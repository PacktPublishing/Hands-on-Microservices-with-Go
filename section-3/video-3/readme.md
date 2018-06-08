# Packt Publishing - Hands on Microservices with Go
# Section 3 - Video 3 - Introduction to gRPC and Protocol Buffers

## Install Protocol Buffers Compiler

```
### Change dir to your Downloads folder
cd ~/Downloads

### Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

### Unzip
unzip protoc-3.3.0-linux-x86_64.zip -d protoc3

### Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

### Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/

### Optional: change owner
sudo chown $USER /usr/local/bin/protoc
sudo chown -R $USER /usr/local/include/google

## Get required go packages

go get google.golang.org/grpc
go get github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go

```


## Generate code

´´´

cd path/to/proto/folder

protoc --go_out=plugins=grpc:. *.proto

´´´


## Learn More

[GRPC Site](https://grpc.io/)

[Google's Protocol Buffers Site](https://developers.google.com/protocol-buffers/)

[Wikipedia - Remote Procedure Call](https://en.wikipedia.org/wiki/Remote_procedure_call)
