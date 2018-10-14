build:
    # 使用如下语句是为了在Windows环境下生成vessel-service/proto/vessel/vessel.pb.go
    # Linux环境中直接部署时无需执行该语句，因为vessel.pb.go已经存在
    # protoc -I. --go_out=plugins=micro:$GOPATH/src/vessel-service proto/vessel/vessel.proto

	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service
