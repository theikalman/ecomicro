package product

import (
	"context"

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
			return nil, errInvalidRequestBody
		}

		return s.CreateProduct(product)
	}
}

func makeGetProductByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		getProductByIDQuery := request.(getProductByIDQuery)
		return s.GetProductByID(getProductByIDQuery.productID)
	}
}

func makeGetProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.GetProducts()
	}
}
