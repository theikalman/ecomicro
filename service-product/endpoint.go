package product

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

func makeGetVersionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Version(), nil
	}
}

func makeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		product, ok := request.(Product)
		if !ok {
			return nil, errors.New("request body decoder is invalid")
		}

		return s.CreateProduct(product)
	}
}

func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetProducts()
	}
}
