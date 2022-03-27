import { Construct } from "constructs"
import { Stack, StackProps } from "aws-cdk-lib"
import { aws_ecr as ecr } from "aws-cdk-lib"

export class RalphbotEcrStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props)

    new ecr.Repository(this, "ralphbotECRRepository", {
      lifecycleRules: [
        {
          maxImageCount: 5,
        },
      ],
      repositoryName: "ralphbot/master",
    })
  }
}
