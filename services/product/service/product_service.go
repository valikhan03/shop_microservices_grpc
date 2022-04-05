package service

import (
	"context"
	"product_service/pb"
	"product_service/repository"
)

type ProductService struct {
	repository repository.Repository
}

func NewProductService(r repository.Repository) *ProductService {
	return &ProductService{
		repository: r,
	}
}


func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	slug, err := s.repository.CreateProduct(req.Title, req.Category, req.Price, req.Description)
	if err != nil{
		return nil, err
	}

	res := &pb.CreateProductResponse{ Product: &pb.Product{
		Slug: slug, 
		Title: req.Title,
		Category: req.Category,
		Price: req.Price,
		Description: req.Description,

	}}

	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	product, err := s.repository.ChangeProduct(req.Product.Slug, req.Product.Title, req.Product.Category, req.Product.Price, req.Product.Description)
	if err != nil{
		return nil, err
	}

	return &pb.UpdateProductResponse{ Product: &pb.Product{
		Slug: product.Slug,
		Title: product.Title,
		Category: product.Category,
		Price: product.Price,
		Description: product.Description,
	}}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error){
	res := &pb.DeleteProductResponse{}
	if s.repository.DeleteProduct(req.Slug){
		res.Result = "DELETED" 
		return res, nil
	}else{
		res.Result = "NOT DELETED"
		return res, nil
	}
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error){
	product, err := s.repository.GetProduct(req.Slug)
	if err != nil{
		return nil, err
	}

	res := &pb.GetProductResponse{Product: &pb.Product{
		Slug: product.Slug,
		Title: product.Title,
		Category: product.Category,
		Price: product.Price,
		Description: product.Description,
	}}

	return res, nil
}

func (s *ProductService) SearchProduct(context.Context, *pb.SearchProductRequest) (*pb.SearchProductResponse, error){

	
	return nil, nil
}





