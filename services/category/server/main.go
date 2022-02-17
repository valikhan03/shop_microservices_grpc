package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"category_service/repository"
	"category_service/service"
	"category_service/pb"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func main() {
	db := initDB()
	repository := repository.NewCategoryRepository(db)
	service := service.NewCategoryService(repository)
	
	grpc_server := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpc_server, service)

	
}

type db_configs struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Database string `yaml:"database"`
	User string `yaml:"user"`
	SSLMode string `yaml:"sslmode"`

}
func initDB() *sqlx.DB{
	var configs db_configs
	data, err := ioutil.ReadFile("../configs/database.yaml")
	if err != nil{
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, configs)
	if err != nil{
		log.Fatal(err)
	}
	godotenv.Load("database.env")

	data_src := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
	configs.Host, configs.Port, configs.User, os.Getenv("DB_PASSWORD"), configs.Database, configs.SSLMode)

	db, err := sqlx.Open("pgx", data_src)
	if err != nil{
		log.Fatal(err)
	}
	return db
}