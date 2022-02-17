package service

import(
	"context"

	"category_service/pb"
	"category_service/repository"
	
)

type CategoryService struct{
	repository repository.Repository
}

func NewCategoryService(r repository.Repository) *CategoryService{
	return &CategoryService{
		repository: r,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error){
	category, err := s.repository.CreateCategory(ctx, req.ParentId, req.Name, req.Path)
	res := &pb.CreateCategoryResponse{
		Category: category,
	}
	return res, err
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error){
	status, err := s.repository.UpdateCategory(ctx, req.GetSlug(), req.GetParentId(), req.GetPath(), req.GetName(), req.GetStatus())
	res := &pb.UpdateCategoryResponse{
		Status: status,
	}
	return res, err
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error){
	status, err := s.repository.DeleteCategory(ctx, req.GetSlug())
	res := &pb.DeleteCategoryResponse{
		Status: status,
	}
	return res, err
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error){
	category, err := s.repository.GetCategory(ctx, req.GetSlug())
	res := &pb.GetCategoryResponse{
		Category: category,
	}
	return res, err
}

func (s *CategoryService) FindCategory(ctx context.Context, req *pb.FindCategoryRequest) (*pb.FindCategoryResponse, error){
	categories, err := s.repository.FindCategory(ctx, req.GetSlug())
	res := &pb.FindCategoryResponse{
		Category: categories,
	}
	return res, err
}
