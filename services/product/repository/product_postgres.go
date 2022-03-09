package repository

import (
	"log"
	"product_service/pb"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
	c *sqlx.Conn
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(name string, category string, price uint32, slug string) (*pb.Product, error) {
	var id int64
	rows, err := r.db.Query("INSERT INTO products (name, category, price, slug) VALUES ($1, $2, $3, $4)", name, category, price, slug)
	if err !=nil{
		return nil, err
	}
	rows.Scan(&id)
	product := &pb.Product{
		Name: name,
		Category: category,
		Price: price,
		Slug: slug,
	}

	return product, nil
}

func (r *ProductRepository) ChangeProduct(id int64, name string, category string, price uint32, slug string) (*pb.Product, error) {
	_, err := r.db.Exec("UPDATE products SET name=$1, category=$2, price=$3, slug=$4 WHERE id=$5", 
		name, category, price, slug, id)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	product := &pb.Product{
		Id: id,
		Name: name,
		Category: category,
		Price: price,
		Slug: slug,
	}

	return product, nil
}

func (r *ProductRepository) DeleteProduct(slug string) bool {
	return false
}

type Product struct{
	Id int64
    Name string
    Category string
    Price uint32 
    Slug string
}

func (r *ProductRepository) GetProduct(slug string) (*pb.Product, error) {
	var product Product
	row, err := r.db.Query("SELECT * FROM products WHERE slug=$1 LIMIT 1", slug)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	row.Scan(&product)
	resultProduct := &pb.Product{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		Category: product.Category,
		Slug: product.Slug,
	}
	
	return resultProduct, nil
}

func (r *ProductRepository) SearchProduct(filter pb.Filter, escape_id int64) (*pb.Product, error) {
	var product Product
	query := "SELECT * FROM products WHERE name LIKE '%'||$1||'%' AND category=$2 AND price<=$3 AND price>=$4 AND id>$5 LIMIT 1"
	row, err := r.db.Query(query, filter.Name, filter.Category, filter.MaxPrice, filter.MinPrice, escape_id)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	row.Scan(&product)
	resultProduct := &pb.Product{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		Category: product.Category,
		Slug: product.Slug,
	}
	
	return resultProduct, nil
}
