FROM golang:latest

RUN apt-get update
RUN apt-get install vim -y
RUN go install "github.com/go-sql-driver/mysql@latest"