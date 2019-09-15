FROM xushikuan/alpine-build:2.0 AS builder

ENV WORK_DIR=$GOPATH/src/github.com/sillyhatxu/mini-mq

WORKDIR $WORK_DIR

COPY . .
#RUN go mod download
#RUN go mod tidy
#RUN go mod vendor

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o main main.go


FROM xushikuan/alpine-build:1.0

ENV BUILDER_WORK_DIR=/go/src/github.com/sillyhatxu/mini-mq
ENV WORK_DIR=/app
ENV TIME_ZONE=Asia/Singapore
WORKDIR $WORK_DIR
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone
RUN mkdir -p logs
RUN mkdir -p db
RUN mkdir -p data

COPY --from=builder $BUILDER_WORK_DIR/main $WORK_DIR
COPY --from=builder $BUILDER_WORK_DIR/config.conf $WORK_DIR
COPY --from=builder $BUILDER_WORK_DIR/db $WORK_DIR/db
COPY --from=builder $BUILDER_WORK_DIR/basic.db $WORK_DIR/data
ENTRYPOINT ./main -c config.conf