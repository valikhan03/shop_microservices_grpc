package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	//"google.golang.org/grpc/credentials/insecure"

	"api_gw/pb/category_service"
)


func main() {
	var router = runtime.NewServeMux()
	
	err := godotenv.Load("services.env")
	if err != nil{
		log.Fatal(err)
	}

	var categoryServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("CATEGORY_SERVICE_HOST"), os.Getenv("CATEGORY_SERVICE_PORT"))


	grpcCategoryServiceConn, err := grpc.Dial(categoryServiceAddr, grpc.WithInsecure())
	fmt.Println("Connected: ", categoryServiceAddr)
	grpcCategoryServiceConn.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer grpcCategoryServiceConn.Close()

	err = category_service.RegisterCategoryServiceHandler(context.Background(), router, grpcCategoryServiceConn)
	if err != nil {
		log.Fatal(err)
	}

	proxy_router := http.NewServeMux()
	proxy_router.HandleFunc("/ping", ping)

	server := &http.Server{
		Addr: "127.0.0.1:8085",
		Handler: http.HandlerFunc(func(resw http.ResponseWriter, req *http.Request) {
			if strings.HasPrefix(req.URL.Path, "/api") {
				router.ServeHTTP(resw, req)
				return
			}
			proxy_router.ServeHTTP(resw, req)
		}),
	}
	fmt.Println("PROXY started working")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func ping(resw http.ResponseWriter, req *http.Request) {
	resw.Write([]byte("WORKS"))
}
