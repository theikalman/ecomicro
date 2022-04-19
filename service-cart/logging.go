package cart

import (
	"time"

	"github.com/go-kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func (s loggingService) Version() string {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "version",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.Version()
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return loggingService{logger: logger, Service: s}
}
