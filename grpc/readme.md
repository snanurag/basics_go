1. Install latest protobuf using `brew install protobuf`
2. go to proto directory. Run `generate_proto.sh`.
3. cd to this path.
4. Run `go get -v`
5. Run `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
6. Run `go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`
7. Run `go get -u github.com/golang/protobuf/protoc-gen-go`
8. Go to grpc folder and run `go run grpc_server.go`
9. Go to entrypoint folder and run `go run entrypoint.go`
10. From web-brower hit `v1/example/hello`. It should work.
11. From any POST web client hit `v1/example/echo`. It should work.