FROM public.ecr.aws/docker/library/golang:1.18.3 as build_image

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify
RUN CGO_ENABLED=0 go build -v -o /tmp/build/ralphbot/ ./...
#CGO_ENABLED=0 to force a static build. Next stage of build uses an image that has libraries in different locations
#https://stackoverflow.com/a/36308464

FROM public.ecr.aws/docker/library/golang:1.18.3-alpine
WORKDIR /app
COPY --from=build_image /tmp/build/ralphbot/ ./
COPY --from=build_image /app/internal/dadjoke/jokes.json ./
CMD [ "./ralphbot" ]