package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type ConsulRegister interface {
	Register(ctx context.Context, instanceId, serviceName, hostAddr string) error
	Deregister(instanceId, serviceName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(instanceId, serviceName string) error
}

func GenerateConsulInstanceId(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
