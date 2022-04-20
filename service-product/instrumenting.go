package product

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (s instrumentingService) Version() Version {
	defer func(begin time.Time) {
		s.requestCount.With("method", "version").Add(1)
		s.requestLatency.With("method", "version").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Version()
}

func (s instrumentingService) CreateProduct(product Product) (Product, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "create").Add(1)
		s.requestLatency.With("method", "create").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.CreateProduct(product)
}

func (s instrumentingService) GetProducts() ([]Product, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "get_products").Add(1)
		s.requestLatency.With("method", "get_products").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetProducts()
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}
