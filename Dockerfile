FROM golang:1.14 as builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN GOOS=linux

RUN go build -o main ./main.go

FROM ubuntu:16.04

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 80

CMD ["/app/main"]