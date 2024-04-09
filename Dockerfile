# building the go service

FROM golang:1.19 as builder
WORKDIR /application/golang-jwt-project
COPY . /application/golang-jwt-project
COPY .env /application/.env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/golang-jwt-project -mod vendor ./cmd/server/main.go

# copying builds to final
FROM alpine:3.10.2 as deploy
WORKDIR /application/

COPY --from=builder /application/golang-jwt-project/build/golang-jwt-project /application/
COPY --from=builder /application/.env /application/

CMD [ "sh", "-c", "/application/golang-jwt-project" ]
