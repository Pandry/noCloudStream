# https://developers.google.com/protocol-buffers/docs/proto3#generating
# protoc --proto_path=IMPORT_PATH --go_out=./auther/api/grpc/ ./auther/api/grpc/auther.proto
protobuf:
	protoc --go_out=. --go-grpc_out=. ./auther/api/grpc/auther.proto 

swagger:
	swagger generate spec -o ./auther/swagger.yaml --scan-models