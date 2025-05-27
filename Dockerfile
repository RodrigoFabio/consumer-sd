FROM golang:1.23

WORKDIR /usr/src/app

COPY . .

RUN go mod init example.com/myapp || true

RUN go mod tidy

RUN go build -v -o app .

CMD ["./app"]
