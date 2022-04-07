package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"product_service/pb"
	"product_service/repository"
	"product_service/service"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func main() {
	db := initDB()
	godotenv.Load("server.env")
	repository := repository.NewProductRepository(db)
	productService := service.NewProductService(repository)

	server := grpc.NewServer()

	pb.RegisterProductServiceServer(server, productService)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started working..")
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

type db_configs struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	SSLMode  string `yaml:"sslmode"`
}

func initDB() *sqlx.DB {
	var configs db_configs
	data, err := ioutil.ReadFile("configs/database.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	data_src := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Host, configs.Port, configs.User, os.Getenv("DB_PASSWORD"), configs.Database, configs.SSLMode)

	db, err := sqlx.Open("pgx", data_src)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
