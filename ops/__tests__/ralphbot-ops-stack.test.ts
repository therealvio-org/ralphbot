import * as cdk from "aws-cdk-lib"
import { AppExtension } from "../src/app"
import { RalphbotStack } from "../src/stacks/ralphbot-ops-stack"
import { Template } from "aws-cdk-lib/assertions"

describe("Retrieve ralphbot ECR Stack", () => {
  const app = new cdk.App()
  const projectName = `ralphbot`

  const commit = "aaa123"
  const environment = "test"
  const version = "2.0.20"
  const accountId = "12345678901234"
  const region = "foo-bar-1"

  const extendedProperties = new AppExtension({
    commit,
    environment,
    projectName,
    version,
  })

  const stack = new RalphbotStack(
    app,
    `${projectName}-ops`,
    {
      env: {
        account: accountId,
        region: region,
      },
    },
    extendedProperties,
  )

  const template = Template.fromStack(stack)

  it("Synths and snapshots", () => {
    expect(template).toMatchSnapshot()
  })
})
