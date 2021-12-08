# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /transaction-service

COPY . /transaction-service
RUN go mod download

WORKDIR /transaction-service/app

RUN go build -o transaction-service
RUN chmod 777 transaction-service

EXPOSE 8081

CMD [ "./transaction-service" ]