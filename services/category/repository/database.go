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

func (r *CategoryRepository) CreateCategory(ctx context.Context, parent_id int32, slug string) (*pb.Category, error){
	
	res, err := r.database.Exec("INSERT INTO category (slug, parent_id) VALUES ($1, $2) RETURNING id", slug, parent_id)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	id, err := res.LastInsertId()
	category := &pb.Category{
		ID: int32(id),
		Slug: slug,
		ParentID: parent_id,
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, id int32, slug string, parent_id int32) (error){
	tx, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil{
		log.Println(err)
		return err
	}

	_, err = tx.Exec("UPADTE category SET slug=$1, parent_id=$2 WHERE id=$3", slug, parent_id, id)
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, slug string) (error){
	tx, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil{
		log.Println(err)
		return err
	}

	_, err = tx.Exec("DELETE FROM category WHERE slug=$1", slug)
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

func (r *CategoryRepository) GetSubCategories(ctx context.Context, slug string) ([]*pb.Category, error){
	return nil, nil
}
