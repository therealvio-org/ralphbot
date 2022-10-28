FROM public.ecr.aws/docker/library/node:18-alpine

WORKDIR /cdk

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .
RUN yarn build

ENTRYPOINT ["yarn"]