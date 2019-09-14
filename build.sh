CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go



docker build --build-arg PROJECT_NAME=mini-mq -t xushikuan/mini-mq .
docker build --build-arg PROJECT_NAME=mini-mq -q -t xushikuan/mini-mq .

docker build -t xushikuan/mini-mq .
docker tag xushikuan/mini-mq:latest xushikuan/mini-mq:1.0
docker push xushikuan/mini-mq:1.0
go clean



docker stack deploy -c docker-compose.yml sillyhat


docker stack deploy -c /root/server/docker-compose.mini-mq.yml sillyhat