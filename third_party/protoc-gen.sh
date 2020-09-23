
protoc --proto_path=${GOPATH}/gateway/proto --proto_path=${GOPATH}/gateway/third_party --go_out=plugins=grpc:${GOPATH}/gateway/proto user-service.proto
protoc --proto_path=${GOPATH}/gateway/proto --proto_path=${GOPATH}/gateway/third_party --grpc-gateway_out=logtostderr=true:${GOPATH}/gateway/proto user-service.proto
protoc --proto_path=${GOPATH}/gateway/proto --proto_path=${GOPATH}/gateway/third_party --swagger_out=logtostderr=true:${GOPATH}/gateway/proto user-service.proto