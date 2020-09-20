# Implementing grpc gateway for handling both requests.

## generating proto files
```
$ protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto user-service.proto
```