## Build image
FROM golang:alpine3.19 AS build

WORKDIR /app
COPY . .

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api .

## Base image
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/api /app/api
COPY config.yaml.docker config.yaml
COPY docs docs
COPY logs/cms/.gitignore logs/cms/.gitignore
COPY logs/web/.gitignore logs/web/.gitignore


# Update the package repository and install curl
RUN apk update && apk add --no-cache curl

EXPOSE 3000
EXPOSE 3001

CMD /app/api