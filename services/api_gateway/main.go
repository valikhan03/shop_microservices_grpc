package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	category_service "api_gw/pb/category_service" 
	product_service "api_gw/pb/product_service"
)

func run() error {
	godotenv.Load("services.env")

	var categoryServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("CATEGORY_SERVICE_HOST"), os.Getenv("CATEGORY_SERVICE_PORT"))
	var productServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("PRODUCT_SERVICE_HOST"), os.Getenv("PRODUCT_SERVICE_PORT"))
	
	
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()



	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	
	err := category_service.RegisterCategoryServiceHandlerFromEndpoint(ctx, mux, categoryServiceAddr, opts)
	if err != nil {
		return err
	}

	err = product_service.RegisterProductServiceHandlerFromEndpoint(ctx, mux, productServiceAddr, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8085", mux)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

/*package main

import (
	"context"
	"fmt"

	//"io/ioutil"
	"log"
	"net/http"
	"os"

	//"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	//"github.com/gin-gonic/gin"

	"google.golang.org/grpc/credentials/insecure"

	"api_category_service/pb/category_service"
)

func main() {
	var router = runtime.NewServeMux()

	err := godotenv.Load("services.env")
	if err != nil {
		log.Fatal(err)
	}

	var categoryServiceAddr = fmt.Sprintf("%s:%s", os.Getenv("CATEGORY_SERVICE_HOST"), os.Getenv("CATEGORY_SERVICE_PORT"))

	/*
	grpcCategoryServiceConn, err := grpc.Dial(categoryServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	grpcCategoryServiceConn.Connect()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected: ", categoryServiceAddr)
	defer grpcCategoryServiceConn.Close()
*/

/*
	//err = category_service.RegisterCategoryServiceHandler(context.Background(), router, grpcCategoryServiceConn)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	category_service.RegisterCategoryServiceHandlerFromEndpoint(context.Background(), router, categoryServiceAddr, opts)
	if err != nil {
		log.Fatal(err)
	}

	//proxy_gin_router := gin.Default()
	//proxy_router := http.NewServeMux()

	/*
		handlerFunc1 := http.HandlerFunc(func(resw http.ResponseWriter, req *http.Request) {
			if strings.HasPrefix(req.URL.Path, "/api") {
				router.ServeHTTP(resw, req)
				return
			}
			proxy_router.ServeHTTP(resw, req)
		})
*/

/*
	handlerFunc2 := http.HandlerFunc(func(resw http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/api") {
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(body))
			router.ServeHTTP(resw, req)
			return
		}
		proxy_router.ServeHTTP(resw, req)
	})*/

/*
	server := &http.Server{
		Addr:    "127.0.0.1:8085",
		Handler: router,
	}
	fmt.Println("PROXY started working")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
*/
