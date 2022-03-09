package repository

import(
	"product_service/pb"
)

type Repository interface{
	CreateProduct(name string, category string, price uint32, slug string) (*pb.Product, error)
	ChangeProduct(id int64, name string, category string, price uint32, slug string) (*pb.Product, error)
	DeleteProduct(slug string) (bool)
	GetProduct(slug string) (*pb.Product, error)
	SearchProduct(filter pb.Filter, escape_id int64) (*pb.Product, error)
}