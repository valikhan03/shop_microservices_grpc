package main

import (
	"net"
	"log"

	"product_service/pb"
	"product_service/service"

	"google.golang.org/grpc"
)

func main() {

	productService := service.NewProductService(nil)

	server := grpc.NewServer(grpc.WithInsecure())

	pb.RegisterProductServiceServer(server, productService)

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil{
		log.Fatal(err)
	}

	err = server.Serve(listener)
	if err != nil{
		panic(err)
	}
}
