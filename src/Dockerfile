FROM public.ecr.aws/docker/library/golang:1.18.3 as build_image

#Build the fargate binary when the "build-deployment-binary" argument is passed, or use "make build" target by default
#This argument shouldn't need to be used in a local development context unless testing the use of the argument
ARG MAKE_TARGET="build"

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify

RUN make ${MAKE_TARGET}

FROM public.ecr.aws/docker/library/golang:1.18.3-alpine
WORKDIR /app
COPY --from=build_image /app/bin/ralphbot ./
COPY --from=build_image /app/internal/dadjoke/jokes.json ./
CMD [ "./ralphbot" ]