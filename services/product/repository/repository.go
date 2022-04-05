package repository

import(
	"product_service/models"
)

type Repository interface{
	CreateProduct(title string, category string, price uint64, description string) (slug string, err error)
	ChangeProduct(slug string, title string, category string, price uint64, description string) (product *models.Product, err error)
	DeleteProduct(slug string) (result bool)
	GetProduct(slug string) (product *models.Product, err error)
	SearchProduct(title string, category string, minprice int64, maxprice int64) (product *models.Product, err error)
}