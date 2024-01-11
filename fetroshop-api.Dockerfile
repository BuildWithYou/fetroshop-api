## Build image
FROM golang:alpine3.19 AS build

WORKDIR /app/api
COPY . .

RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fetroshop-api .

## Base image
FROM alpine:latest

WORKDIR /app/api
COPY --from=build /app/api/fetroshop-api /app/api/fetroshop-api
COPY config.yaml.docker config.yaml
COPY docs docs


# Update the package repository and install curl
RUN apk update && apk add --no-cache curl

EXPOSE 3000
EXPOSE 3001

CMD /app/api/fetroshop-api