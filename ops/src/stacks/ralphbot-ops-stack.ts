import { AppExtensionProps } from "../app"
import { Construct } from "constructs"
import { Stack, StackProps } from "aws-cdk-lib"
import { aws_ec2 as ec2 } from "aws-cdk-lib"
import { aws_ecr as ecr } from "aws-cdk-lib"
import { aws_ecs as ecs } from "aws-cdk-lib"
import { aws_iam as iam } from "aws-cdk-lib"
import { aws_logs as logs } from "aws-cdk-lib"
import { aws_secretsmanager as secretsmanager } from "aws-cdk-lib"
import { aws_ssm as ssm } from "aws-cdk-lib"

export class RalphbotStack extends Stack {
  constructor(
    scope: Construct,
    id: string,
    props?: StackProps,
    extendedProps?: AppExtensionProps
  ) {
    super(scope, id, props)

    const vpc = new ec2.Vpc(this, "vpc", {
      maxAzs: 1,
      natGateways: 0
    })

    const repositoryRef = ecr.Repository.fromRepositoryArn(
      this,
      "ralphbotECRRepositoryRef",
      `arn:aws:ecr:${props?.env?.region}:${props?.env?.account}:repository/ralphbot/master`
    )

    const ralphbotImage = ecs.ContainerImage.fromEcrRepository(
      repositoryRef,
      extendedProps?.version
    )

    const cluster = new ecs.Cluster(this, "Cluster", {
      vpc: vpc,
    })
    const taskDefinition = new ecs.FargateTaskDefinition(
      this,
      "TaskDefinition",
      {
        cpu: 256,
        memoryLimitMiB: 512,
      }
    )

    const botToken = secretsmanager.Secret.fromSecretNameV2(
      this,
      "botTokenFromName",
      "ralphbot/token"
    )
    const secretArnSuffix = ssm.StringParameter.valueForStringParameter(
      this,
      "/ralphbot/token/arn-suffix"
    )

    taskDefinition.addToExecutionRolePolicy(
      new iam.PolicyStatement({
        actions: [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret",
        ],
        resources: [`${botToken.secretArn}-${secretArnSuffix}`],
      })
    )

    taskDefinition.addContainer("ralphbotContainer", {
      image: ralphbotImage,
      secrets: {
        BOT_TOKEN: ecs.Secret.fromSecretsManager(botToken, "token"),
      },
      logging: ecs.LogDriver.awsLogs({
        streamPrefix: "ralphbot",
        logRetention: logs.RetentionDays.ONE_WEEK,
      }),
    })
    new ecs.FargateService(this, "Service", {
      assignPublicIp: true,
      cluster: cluster,
      taskDefinition: taskDefinition,
      circuitBreaker: { rollback: true },
      vpcSubnets: vpc.selectSubnets({
        subnetType: ec2.SubnetType.PUBLIC,
      }),
    })
  }
}
