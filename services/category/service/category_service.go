package service

import (
	"category_service/pb"
	"category_service/repository"
	"context"
	"fmt"
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
	category, err := s.repository.CreateCategory(ctx, req.GetParentID(), req.GetName())
	if err != nil{
		return nil, err
	}

	res := &pb.CreateCategoryResponse{
		Category: category,
	}

	return res, nil
}

func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error){
	err := s.repository.UpdateCategory(ctx, req.Category.ID, req.Category.Name, req.Category.ParentID)
	if err != nil{
		return nil, err
	}

	res := &pb.UpdateCategoryResponse{Category: req.Category}
	return res, nil
}

func (s *CategoryService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error){
	err := s.repository.DeleteCategory(ctx, req.GetName())
	if err != nil{
		return nil, err
	}

	res := &pb.DeleteCategoryResponse{ Result: fmt.Sprintf("Category %s deleted", req.GetName())}
	return res, err
}

func (s *CategoryService) GetAllCategories(ctx context.Context, req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error){
	categories, err := s.repository.GetAllCategories(ctx)
	if err != nil{
		return nil, err
	}
	res := &pb.GetAllCategoriesResponse{
		Category: categories,
	}

	return res, nil
}

func (s *CategoryService) GetSubCategories(ctx context.Context, req *pb.GetSubCategoriesRequest) (*pb.GetSubCategoriesResponse, error){
	categories, err := s.repository.GetSubCategories(ctx, req.GetName())
	if err != nil{
		return nil, err
	}

	res := &pb.GetSubCategoriesResponse{ Category: categories }
	return res, err
}