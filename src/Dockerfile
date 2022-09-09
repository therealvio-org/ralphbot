FROM public.ecr.aws/docker/library/golang:1.18.3 as build_image

WORKDIR /app

COPY . ./

RUN go mod download && go mod verify
RUN GOOS=linux CGO_ENABLED=0 go build -v -o bin/ralphbot ./cmd/ralphbot

FROM public.ecr.aws/docker/library/golang:1.18.3-alpine
WORKDIR /app
COPY --from=build_image /app/bin/ralphbot ./
COPY --from=build_image /app/internal/dadjoke/jokes.json ./
CMD [ "./ralphbot" ]