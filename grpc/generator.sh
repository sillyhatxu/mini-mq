cd consumer
protoc --go_out=plugins=grpc:. *.proto
cd ..
cd health
protoc --go_out=plugins=grpc:. *.proto
cd ..
cd producer
protoc --go_out=plugins=grpc:. *.proto
cd ..
cd register
protoc --go_out=plugins=grpc:. *.proto
cd ..
