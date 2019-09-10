FROM golang:1.13 AS builder

ENV GOPATH=/usr/local/go/src/github.com/sillyhatxu
#ARG PROJECT_NAME
ENV PROJECT_NAME=mini-mq
WORKDIR $GOPATH/$PROJECT_NAME
#ADD . $GOPATH/$PROJECT_NAME
COPY . $GOPATH/$PROJECT_NAME

RUN apt-get update && apt-get install -y gcc-aarch64-linux-gnu
RUN CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=amd64 go build -o main main.go

FROM xushikuan/alpine-build:1.0

ENV GOPATH=/usr/local/go/src/github.com/sillyhatxu
ENV PROJECT_NAME=mini-mq
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR /go
COPY --from=build /app/bin/server /app/bin/server
COPY --from=builder $GOPATH/$PROJECT_NAME/main .
COPY --from=builder $GOPATH/$PROJECT_NAME/config.conf .
COPY --from=builder $GOPATH/$PROJECT_NAME/db ./data
ENTRYPOINT ./main -c config.conf

#FROM xushikuan/alpine-build:1.0
#
#ENV GOPATH=/usr/local/go/src/github.com/sillyhatxu
#ARG PROJECT_NAME
#ENV TIME_ZONE=Asia/Singapore
#RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
#
#WORKDIR /go
#COPY main .
#COPY db ./data
#COPY config.conf .
#ENTRYPOINT ./main -c config.conf


#RUN               GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main main.go
#RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main main.go

#FROM golang:1.13-alpine
#
#ENV GOPATH=/go
#ENV TIME_ZONE=Asia/Singapore
#
#RUN apk add --no-cache tzdata
#RUN apk --update --no-cache add curl
#RUN apk add --no-cache ca-certificates
#RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
#RUN CGO_ENABLED=1 GOOS=linux go install -a server
#
#WORKDIR /go
#COPY main .
#COPY db ./data
#COPY config.conf .
#ENTRYPOINT ./main -c config.conf