package repository

import(
	"context"
	"category_service/pb"
)

type Repository interface{
	CreateCategory(ctx context.Context, parent_id string, name string, path string) (*pb.Category, error)
	UpdateCategory(ctx context.Context, slug string, parent_id string, path string, name string, status int32) (int32, error)
	DeleteCategory(ctx context.Context, slug string) (int32, error)
	GetCategory(ctx context.Context, slug string) (*pb.Category, error)
	FindCategory(ctx context.Context, slug string) ([]*pb.Category, error)
}