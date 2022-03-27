import { App as BaseApp, AppProps as BaseAppProps } from "aws-cdk-lib"

export interface AppExtensionProps extends BaseAppProps {
  projectName: string
  commit: string
  environment: string
  version: string
}

export class AppExtension extends BaseApp {
  public readonly commit: string
  public readonly environment: string
  public readonly projectName: string
  public readonly version: string

  constructor(props: AppExtensionProps) {
    super(props)

    this.projectName = props.projectName
    this.commit = props.commit
    this.environment = props.environment
    this.version = props.version
  }
}
