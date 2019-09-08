CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main main.go
docker build --build-arg PROJECT_NAME=mini-mq -t xushikuan/mini-mq .

docker build -t xushikuan/mini-mq .
docker tag xushikuan/mini-mq:latest xushikuan/mini-mq:1.0
docker push xushikuan/mini-mq:1.0
go clean