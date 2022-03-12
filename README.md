# ralphbot

A humble discord bot

## Tools

For programming language, and other tools, check out `.tool-versions` (I tend to use asdf where I can!)

- asdf version manager: https://asdf-vm.com/guide/getting-started.html#getting-started
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

Assuming you have Docker (or any appropriate containerisation tool) installed, and ready to go, the commands are pretty trivial:

### Build the Image locally

```sh
docker build \
-t "<name-for-container>:<tag>"
#e.g. docker build -t "ralphbot-test:unstable"
```

### Run the container

```sh
docker run -it \
-e BOT_TOKEN \
<name-for-container>:<tag>
```

_Note_: BOT_TOKEN needs to be defined.
