import * as ecr from "aws-cdk-lib/aws-ecr"
import { Construct } from "constructs"
import { Stack, StackProps } from "aws-cdk-lib/core"

export class RalphbotEcrStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props)

    new ecr.Repository(this, "ralphbotECRRepository", {
      lifecycleRules: [
        {
          maxImageCount: 10,
        },
      ],
      repositoryName: "ralphbot/master",
    })
  }
}
