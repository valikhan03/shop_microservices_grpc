package repository

import (
	"log"
	"strings"
	"product_service/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(title string, category string, price uint64, description string) (slug string, err error) {
	
	slug = generateSlug(title) 
	query := "INSERT INTO products (slug, title, price, description, category_id) VALUES ($1, $2, $3, $4, SELECT id FROM category WHERE slug=$5)"
	_, err = r.db.Exec(query, slug, title, price, description, category)
	if err != nil{
		log.Println(err)
		return "", err
	}

	return slug, nil

}

func (r *ProductRepository) ChangeProduct(slug string, title string, category string, price uint64, description string) (product *models.Product, err error) {
	query := "UPDATE products (title, category_id, price, description) SET ()"
	_, err = r.db.Exec(query, title, category, price, description)
	if err != nil{
		log.Println(err)
		return
	}
	product = &models.Product{
		Slug: slug,
		Title: title,
		Category: "",
		Price: price,
		Description: description,
	}

	return
}

func (r *ProductRepository) DeleteProduct(slug string) (result bool) {
	result = false
	_, err := r.db.Exec("DELETE from products WHERE slug=$1", slug)
	if err != nil{
		log.Println(err)
		return
	}
	result = true
	return
}

func (r *ProductRepository) GetProduct(slug string) (product *models.Product, err error) {
	return
}
func (r *ProductRepository) SearchProduct(title string, category string, minprice int64, maxprice int64) (product *models.Product, err error){
	return
}


/*

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
*/


func generateSlug(title string) string {
	bannedSymbols := []string{".", ",", "/", "\\", "|", "?", "=", "+", "*"}
	for _, s := range bannedSymbols{
		if strings.Contains(title, s){
			strings.Replace(title, s, "", -1)
		}
	}

	slug := strings.Replace(strings.ToLower(title), " ", "_", -1)

	return slug
}