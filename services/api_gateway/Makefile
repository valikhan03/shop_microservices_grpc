protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb --grpc-gateway_out=:pb

protoc -I . --grpc-gateway_out ./gen/go  --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative  proto/category_service.proto

protoc --proto_path=proto proto/*.proto  --go_out=:. --go-grpc_out=:. --grpc-gateway_out=. --openapiv2_out=:swagger