package repository

import(
	"context"
	"category_service/pb"
)

type Repository interface{
	CreateCategory(ctx context.Context, parent_id int32, name string) (*pb.Category, error)
	UpdateCategory(ctx context.Context, id int32, name string, parent_id int32) (error)
	DeleteCategory(ctx context.Context, name string) (error)
	GetAllCategories(ctx context.Context) ([]*pb.Category, error)
	GetSubCategories(ctx context.Context, name string) ([]*pb.Category, error)
}