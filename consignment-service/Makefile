build:
    # 使用如下语句是为了在Windows环境下生成consignment-service/proto/consignment/consignment.pb.go
    # Linux环境中直接部署时无需执行该语句，因为consignment.pb.go已经存在
    # protoc -I. --go_out=plugins=micro:$GOPATH/src/consignment-service proto/consignment/consignment.proto

    GOOS=linux GOARCH=amd64 go build
    docker build -t consignment-service .

run:
    docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service