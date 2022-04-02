#!/usr/bin/env node
import "source-map-support/register"
import * as cdk from "aws-cdk-lib"
import { AppExtension } from "./app"
import { RalphbotEcrStack } from "./stacks/ralphbot-ecr-stack"
import { RalphbotStack } from "./stacks/ralphbot-ops-stack"

const {
  COMMIT: commit,
  ENVIRONMENT: environment,
  VERSION: version,
  AWS_ACCOUNT_ID: accountId,
  REGION: region,
} = process.env

const projectName = `ralphbot`

if (!environment) {
  throw new Error("ENVIRONMENT environment variable must be supplied")
}

if (!commit) {
  throw new Error("COMMIT environment variable must be supplied")
}

if (!version) {
  throw new Error("VERSION environment variable must be supplied")
}

const extendedProperties = new AppExtension({
  commit,
  environment,
  projectName,
  version,
})

const app = new cdk.App()
new RalphbotEcrStack(app, `${projectName}-ecr`)
new RalphbotStack(
  app,
  `${projectName}-ops`,
  {
    env: {
      account: accountId,
      region: region,
    },
  },
  extendedProperties
)
