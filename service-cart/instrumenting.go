package cart

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func (s instrumentingService) Version() string {
	defer func(begin time.Time) {
		s.requestCount.With("method", "version").Add(1)
		s.requestLatency.With("method", "version").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.Version()
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}
