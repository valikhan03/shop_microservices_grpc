package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"category_service/pb"
	"category_service/repository"
	"category_service/service"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func main() {
	db := initDB()
	godotenv.Load("server.env")
	repository := repository.NewCategoryRepository(db)
	service := service.NewCategoryService(repository)

	grpc_server := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpc_server, service)
	listener, err := net.Listen("tcp", "127.0.0.1:8091")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started working..")
	err = grpc_server.Serve(listener)
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
	godotenv.Load("database.env")

	data_src := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Host, configs.Port, configs.User, os.Getenv("DB_PASSWORD"), configs.Database, configs.SSLMode)

	db, err := sqlx.Open("pgx", data_src)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
