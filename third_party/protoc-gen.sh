protoc --proto_path=proto --proto_path=${GOPATH}/gateway/third_party/google/protobuf --go_out=plugins=grpc:proto user-service.proto