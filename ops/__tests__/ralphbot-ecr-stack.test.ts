import * as cdk from "aws-cdk-lib/core"
import { RalphbotEcrStack } from "../src/stacks/ralphbot-ecr-stack"
import { Template } from "aws-cdk-lib/assertions"

describe("Retrieve ralphbot ECR Stack", () => {
  const app = new cdk.App()
  const projectName = `ralphbot`

  const stack = new RalphbotEcrStack(app, `${projectName}-ecr`)

  const template = Template.fromStack(stack)

  it("Synths and snapshots", () => {
    expect(template).toMatchSnapshot()
  })
})
