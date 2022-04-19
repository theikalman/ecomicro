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
