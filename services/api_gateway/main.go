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

	"github.com/gin-gonic/gin"

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
	if err != nil{
		log.Fatal(err)
	}
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

	proxy_gin_router := gin.Default()

	proxy_router := http.NewServeMux()
	proxy_router.HandleFunc("/ping", ping)

/*
	handlerFunc1 := http.HandlerFunc(func(resw http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/api") {
			router.ServeHTTP(resw, req)
			return
		}
		proxy_router.ServeHTTP(resw, req)
	})
*/
	handlerFunc2 := http.HandlerFunc(func(resw http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/api") {
			router.ServeHTTP(resw, req)
			return
		}
		proxy_gin_router.ServeHTTP(resw, req)
	})

	server := &http.Server{
		Addr: "127.0.0.1:8085",
		Handler: handlerFunc2,
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
