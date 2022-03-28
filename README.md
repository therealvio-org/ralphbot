# ralphbot

A humble discord bot

## Tools

For programming language, and other tools, check out `.tool-versions` (I tend to use asdf where I can!)

- asdf version manager: https://asdf-vm.com/guide/getting-started.html#getting-started
- AWS CLI v2: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-version.html
- Docker: https://docs.docker.com/get-docker/

## Project Structure

- Golang Standards: https://github.com/golang-standards/project-layout
- How do I structure my go project?: https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/

## Use of BOT_TOKEN

Within the source code, you may have noticed that BOT_TOKEN is retrieving an environment variable. Personally, I use `direnv` (read more [here](https://direnv.net/)) installed via `asdf` and place a `.envrc` file in the root of the repository and work from that.

When deployed in "production", because the bot is running as a containerised app, the environment variable can just be passed through a secrets manager tool, like in [AWS ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-envvar).

All in all, this makes local development portable, and doesn't require production-specific logic to be put in place.

Ultimately it is up to you how you want to define the environment variable, though you need to make sure not to leak any tokens in source code :^)

## Working with the container

Assuming you have Docker (or any appropriate containerisation tool) installed, and your `$(pwd)` is `src/`, the basic commands are:

### Build the Image locally

```sh
docker build \
-t "<name-for-container>:<tag>" .
#e.g. docker build -t "ralphbot-test:unstable"
```

### Run the container

```sh
docker run -it \
-e BOT_TOKEN \
<name-for-container>:<tag>
```

_Notes_:

- BOT_TOKEN **needs** to be defined.
- GUILD_ID can be passed optionally, if you wish to develop slash-commands without making them global by default. Refer to: https://discord.com/developers/docs/interactions/application-commands

## Infrastructure

Ralphbot's infrastructure stack is defined in `./ops`. This is an AWS CDK Stack.

If you intend on forking, or cloning this project for your own use, you do not necessarily need to use the CDK Stack, and can define your own infrastructure stack as you please (that's why containers are awesome!).

### Handling BOT_TOKEN as an AWS Secrets Manager secret

It is important to note, that there are some limitations with AWS CDK when it comes to interacting with Secrets manager:

When referencing the ARN of a secret in the CDK stack, CDK deliberately does not insert the trialing characters, and in some scenarios it substitutes them as: `-?????`, requiring a means to update the reference where appropriate. To get around this, you can "predict" the ARN by constructing it, or reference appropriate properties and insert the suffix.

However, because this is a public repository, revealing details such as this in the open would be a bad idea :stuck_out_tongue_closed_eyes:. So, an AWS Systems Manager Parameter is used to fulfil this, so that during the deploy of the CDK stack, the Parameter value is inserted.

#### Deploying the Secret and Parameter Store value

If this is the first time the stack needs to be deployed (or is being redeployed), assuming you have AWS CLI installed, and a valid AWS exec session running, before attempting to deploy the CDK stack, you must:

**IMPORTANT**: Make sure to **NOT** commit changes of the variable definitions in the shell script, or else you risk leaking secrets when you perform the steps below. It's recommended you do these changes in a throw-away git branch, and then delete the branch when you're done.

1. Open the `./ops/bin/create-bot-token-secret.sh` file, and modify the variables appropriately, then save your changes.
2. Run the script: `./ops/bin/create-bot-token-secret.sh`
3. Log in to the AWS console the secrets have been deployed to
4. Locate the Secret in the `Secrets Manager` console
5. Make note of the trialing characters of the Secret ARN. Denoted by `-abcde`, copy this value.
6. Open the `./ops/bin/create-secret-suffix-parameter.sh` file, and modify the variables appropriately, then save your changes.
   - This is the part where you paste the `-abcde` string into the appropriate variable
7. Run the script: `./ops/bin/create-secret-suffix-parameter.sh`

With this completed, the Bot Token is deployed as an AWS Secrets Manager Secret, and the suffix to the Secret ARN is an AWS Systems Manager Parameter. And no secrets are leaked in plaintext in source code :grin:
