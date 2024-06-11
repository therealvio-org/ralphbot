import * as ec2 from "aws-cdk-lib/aws-ec2"
import * as ecr from "aws-cdk-lib/aws-ecr"
import * as ecs from "aws-cdk-lib/aws-ecs"
import * as iam from "aws-cdk-lib/aws-iam"
import * as logs from "aws-cdk-lib/aws-logs"
import * as secretsmanager from "aws-cdk-lib/aws-secretsmanager"
import * as ssm from "aws-cdk-lib/aws-ssm"
import { AppExtensionProps } from "../app"
import { Construct } from "constructs"
import { Stack, StackProps } from "aws-cdk-lib/core"

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
      natGateways: 0,
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
        runtimePlatform: {
          cpuArchitecture: ecs.CpuArchitecture.ARM64,
        },
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
      environment: {
        // trigger removal of registered commands on bot startup
        // this is configured to true on the app code itself, we're surfacing it here so it remains unobscured
        REMOVE_COMMANDS: "true",
      },
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
