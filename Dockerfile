FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .


RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ADD . .

EXPOSE 9993