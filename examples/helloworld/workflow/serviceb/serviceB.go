package serviceb

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type ServiceB interface {
	Join(ctx context.Context, a1 string, a2 string) (string, error)
}

type ServiceBImpl struct {
	cache backend.Cache
}

func NewServiceB(ctx context.Context, cache backend.Cache) (*ServiceBImpl, error) {
	return &ServiceBImpl{cache}, nil
}

func (s *ServiceBImpl) Join(ctx context.Context, a1 string, a2 string) (string, error) {
	return a1 + a2, nil
}
