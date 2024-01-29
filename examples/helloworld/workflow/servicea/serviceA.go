// Package servicea implements a simple Service called ServiceA that calls ServiceB
package servicea

import (
	"context"

	"github.com/blueprint-uservices/tutorial/examples/helloworld/workflow/serviceb"
)

// ServiceA provides the world-facing interface for service A
type ServiceA interface {
	Hello(ctx context.Context, arg string) (string, error)
}

type ServiceAImpl struct {
	serviceB serviceb.ServiceB
}

func NewServiceA(ctx context.Context, serviceB serviceb.ServiceB) (*ServiceAImpl, error) {
	return &ServiceAImpl{serviceB}, nil
}

func (s *ServiceAImpl) Hello(ctx context.Context, arg string) (string, error) {
	return s.serviceB.Join(ctx, "Hello, ", arg)
}
