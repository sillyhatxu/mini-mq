FROM golang:1.12 AS builder

ENV GOPATH=/usr/local/go/src/github.com/sillyhatxu
ARG PROJECT_NAME
WORKDIR $GOPATH/$PROJECT_NAME
COPY . $GOPATH/$PROJECT_NAME

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM xushikuan/alpine-build:1.0

ENV GOPATH=/go
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR $GOPATH
COPY --from=builder $GOPATH/$PROJECT_NAME/$MODULE_NAME/main /go
COPY ./main $GOPATH
COPY ./db $GOPATH
COPY ./config.conf $GOPATH
#RUN mkdir -p logs
#RUN mkdir -p data
#ENTRYPOINT ./main -c config.conf