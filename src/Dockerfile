FROM public.ecr.aws/docker/library/golang:1.17.8 as build_image

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify
RUN CGO_ENABLED=0 go build -v -o /tmp/build/ralphbot ./...
#CGO_ENABLED=0 to force a static build. Next stage of build uses an image that doesn't have *all* libraries

FROM public.ecr.aws/docker/library/golang:1.17.8-alpine
WORKDIR /app
COPY --from=build_image /tmp/build/ralphbot ./
CMD [ "./ralphbot" ]