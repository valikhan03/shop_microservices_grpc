package service

import(
	"context"
	"product_service/pb"
	"product_service/repository"
)

type ProductService struct{
	repository repository.Repository
}



func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error){
	product, err := s.repository.CreateProduct(req.Name, req.Category, req.Price, req.Slug)
	if err != nil{
		return nil, err
	}

	res := &pb.CreateProductResponse{
		Product: product,
	}

	return res, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error){
	product, err := s.repository.ChangeProduct(req.Product.Id, req.Product.Name, req.Product.Category, req.Product.Price, req.Product.Slug)
	if err != nil{
		return nil, err
	}

	res := &pb.UpdateProductResponse{
		Product: product,
	}

	return res, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error){
	status := s.repository.DeleteProduct(req.Slug)
	res := &pb.DeleteProductResponse{
		Status: status,
	}

	return res, nil
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error){
	product, err := s.repository.GetProduct(req.Slug)
	if err != nil{
		return nil, err
	}

	res := &pb.GetProductResponse{
		Product: product,
	}

	return res, nil
}

func (s *ProductService) SearchProduct(req *pb.SearchProductRequest, stream pb.ProductService_SearchProductServer) error{
	var last_id int64 = 0
	for{
		product, err := s.repository.SearchProduct(*req.GetFilter(), last_id)
		if err != nil{
			return err
		}
		last_id = product.Id
		err = stream.Send(&pb.SearchProductResponse{ Product: product})
		if err != nil{
			return err
		}
	}
}
