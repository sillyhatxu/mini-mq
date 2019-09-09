FROM golang:1.12 AS builder

ENV GOPATH=/usr/local/go/src/github.com/sillyhatxu
ARG PROJECT_NAME
WORKDIR $GOPATH/$PROJECT_NAME
COPY . $GOPATH/$PROJECT_NAME

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM xushikuan/alpine-build:1.0

ENV GOPATH=/usr/local/go/src/github.com/sillyhatxu
ARG PROJECT_NAME
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR /go
COPY --from=builder $GOPATH/$PROJECT_NAME/main .
COPY --from=builder $GOPATH/$PROJECT_NAME/config.conf .
COPY --from=builder $GOPATH/$PROJECT_NAME/db ./data
ENTRYPOINT ./main -c config.conf