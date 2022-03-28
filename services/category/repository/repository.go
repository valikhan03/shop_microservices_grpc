package repository

import(
	"context"
	"category_service/pb"
)

type Repository interface{
	CreateCategory(ctx context.Context, parent_id int32, slug string) (*pb.Category, error)
	UpdateCategory(ctx context.Context, id int32, slug string, parent_id int32) (error)
	DeleteCategory(ctx context.Context, slug string) (error)
	GetAllCategories(ctx context.Context) ([]*pb.Category, error)
	GetSubCategories(ctx context.Context, slug string) ([]*pb.Category, error)
}