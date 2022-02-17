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

func (r *CategoryRepository) CreateCategory(ctx context.Context, parent_id string, name string, path string) (*pb.Category, error) {
	var category = new(pb.Category)
	tx, err := r.database.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	_, err = tx.ExecContext(ctx, "insert into  category () values ()")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, slug string, parent_id string, path string, name string, status int32) (int32, error) {
	tx, err := r.database.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	_, err = tx.Exec("update")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return 1, nil
}

func (r *CategoryRepository) DeleteCategory(ctx context.Context, slug string) (int32, error) {
	tx, err := r.database.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	_, err = tx.Exec("delete from category where slug=$1", slug)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return 1, err
}

func (r *CategoryRepository) GetCategory(ctx context.Context, slug string) (*pb.Category, error) {
	var category = new(pb.Category)
	err := r.database.Get(category, "select * from category where slug=$1", slug)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) FindCategory(ctx context.Context, slug string) ([]*pb.Category, error) {
	var categories []*pb.Category
	err := r.database.Select(categories, "select * from category where slug like %$1%", slug)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return categories, err
}
