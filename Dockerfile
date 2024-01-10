## Build image
FROM golang:alpine3.19 AS build

WORKDIR /app/fetroshop-api
COPY . .

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fetroshop-api .

## Base image
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/fetroshop-api/fetroshop-api /app/api
COPY config.yaml.docker config.yaml
COPY docs docs


# Update the package repository and install curl
RUN apk update && apk add --no-cache curl

EXPOSE 3000
EXPOSE 3001

CMD /app/api