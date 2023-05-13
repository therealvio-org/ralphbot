FROM public.ecr.aws/docker/library/node:18.16-alpine

WORKDIR /cdk

COPY package.json pnpm-lock.yaml ./

# pnpm isn't installed by default
# https://pnpm.io/installation#using-corepack
RUN corepack enable

RUN pnpm install --frozen-lockfile

COPY . .

ENTRYPOINT ["pnpm"]