#FROM xushikuan/alpine-build:1.0
FROM golang:1.13-alpine

ENV WORK_DIR=/go
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR $WORK_DIR
RUN mkdir -p logs
COPY main .
COPY config.conf .
COPY basic.db ./data
COPY db db
RUN ls
RUN pwd

ENTRYPOINT ./main -c config.conf