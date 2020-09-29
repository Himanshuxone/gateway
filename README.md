# Implementing grpc gateway for handling both requests.

## generating proto files
```
$ protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto user-service.proto
```

### Rest service access:-

#### Create User: Method: POST
Request:-
```
{
    "id":123,
    "user": "Irene A. Andrews",
    "email": "IreneAAndrews@teleworm.us"
}
```
Response:-
```
{
    "response": {
        "id": "12345",
        "user": "Irene A. Andrews"
    }
}
```