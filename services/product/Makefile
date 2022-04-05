create_migrations: migrate create -ext sql -dir schema -seq product_service_db
gen_proto: protoc -I ./protobuf --go_out=. --go-grpc_out=. protobuf/product_service.proto