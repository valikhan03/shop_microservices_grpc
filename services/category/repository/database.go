package repository

import (
	"category_service/pb"
	"context"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)


type CategoryRepository struct {
	database *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		database: db,
	}
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, parent_id int32, name string) (*pb.Category, error){
	tx, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil{
		log.Println(err)
		return nil, err
	}

	rows, err := tx.Query("INSERT INTO category (name, parent_id) VALUES ($1, $2) RETURNING id", name, parent_id)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	var id int32
	err = rows.Scan(&id)
	if err != nil{
		log.Fatal(err)
	}

	category := &pb.Category{
		ID: id,
		Name: name,
		ParentID: parent_id,
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id int32, name string, parent_id int32) (error){
	tx, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil{
		log.Println(err)
		return err
	}

	_, err = tx.Exec("UPADTE category SET name=$1, parent_id=$2 WHERE id=$3", name, parent_id, id)
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, name string) (error){
	tx, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil{
		log.Println(err)
		return err
	}

	_, err = tx.Exec("DELETE FROM category WHERE name=$1", name)
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func (r *CategoryRepository) GetAllCategories(ctx context.Context) ([]*pb.Category, error){
	tx, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil{
		log.Println(err)
		return nil, err
	}	
	var categories []*pb.Category
	err = tx.Select(&categories, "SELECT * FROM category WHERE parent_id=NULL")
	

	return nil, nil
}

func (r *CategoryRepository) GetSubCategories(ctx context.Context, name string) ([]*pb.Category, error){
	return nil, nil
}
