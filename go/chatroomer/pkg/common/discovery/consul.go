package discovery

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

type ConsulRegistry struct {
	client *consul.Client
}

func NewConsulRegistry(addr, serviceName string) (*ConsulRegistry, error) {
	config := consul.DefaultConfig()
	config.Address = addr

	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConsulRegistry{
		client: client,
	}, nil
}

func (r *ConsulRegistry) Register(
	ctx context.Context,
	instanceId, serviceName, hostAddr string,
) error {
	parts := strings.Split(hostAddr, ":")
	if len(parts) != 2 {
		return errors.New("invalid host:port format")
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	host := parts[0]

	ags := &consul.AgentServiceCheck{
		CheckID:                        instanceId,
		TLSSkipVerify:                  true,
		TTL:                            "5s",
		Timeout:                        "1s",
		DeregisterCriticalServiceAfter: "10s",
	}

	asr := &consul.AgentServiceRegistration{
		ID:      instanceId,
		Address: host,
		Port:    port,
		Name:    serviceName,
		Check:   ags,
	}

	return r.client.Agent().ServiceRegister(asr)
}

func (r *ConsulRegistry) Deregister(instanceId, serviceName string) error {
	return r.client.Agent().CheckDeregister(instanceId)
}

func (r *ConsulRegistry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var instances []string
	for _, entry := range entries {
		instances = append(
			instances,
			fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port),
		)
	}
	return instances, nil
}

func (r *ConsulRegistry) HealthCheck(instanceId, serviceName string) error {
	return r.client.Agent().UpdateTTL(instanceId, "online", consul.HealthPassing)
}
