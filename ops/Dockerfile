FROM public.ecr.aws/docker/library/node:20.0-alpine

WORKDIR /cdk

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .
RUN yarn build

ENTRYPOINT ["yarn"]