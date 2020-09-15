FROM golang:latest

WORKDIR /calculator

COPY . .

RUN go build

EXPOSE 8080

ENTRYPOINT ["./calculator","config","config.properties"]

