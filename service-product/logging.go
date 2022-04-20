package product

import (
	"time"

	"github.com/go-kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func (s loggingService) Version() Version {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "version",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Version()
}

func (s loggingService) CreateProduct(product Product) (Product, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.CreateProduct(product)
}

func (s loggingService) GetProducts() ([]Product, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get_products",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetProducts()
}

func (s loggingService) GetProductByID(productID uint) (Product, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "get_by_id",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetProductByID(productID)
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return loggingService{logger: logger, Service: s}
}
