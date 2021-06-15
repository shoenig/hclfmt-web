package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
)

var (
	ErrNoBindAddress = errors.New("no bind address")
	ErrNoBindPort    = errors.New("no bind port")
)

const (
	defaultReadTimeout  = 3 * time.Second
	defaultWriteTimeout = 3 * time.Second
)

type Configuration struct {
	Service       string `hcl:"service"`
	BindAddress   string `hcl:"bind_address"`
	BindPort      int    `hcl:"bind_port"`
	MaxRequestLen int    `hcl:"max_request_len"`
}

func logEnvironment(name string) {
	value := os.Getenv(name)
	if value == "" {
		value = "<unset>"
	}
	log.Printf("environment %s = %s", name, value)
}

func Consul() (*consulapi.Client, error) {
	logEnvironment("CONSUL_HTTP_ADDR")
	logEnvironment("CONSUL_NAMESPACE")
	logEnvironment("CONSUL_CACERT")
	logEnvironment("CONSUL_CLIENT_CERT")
	logEnvironment("CONSUL_CLIENT_KEY")
	logEnvironment("CONSUL_HTTP_SSL")
	logEnvironment("CONSUL_HTTP_SSL_VERIFY")
	logEnvironment("CONSUL_TLS_SERVER_NAME")
	logEnvironment("CONSUL_GRPC_ADDR")
	logEnvironment("CONSUL_HTTP_TOKEN_FILE")
	consulConfig := consulapi.DefaultConfig()
	return consulapi.NewClient(consulConfig)
}

func (c Configuration) Address() string {
	return fmt.Sprintf("%s:%d", c.BindAddress, c.BindPort)
}

func (c Configuration) GetService() (*connect.Service, error) {
	if c.BindAddress == "" {
		return nil, ErrNoBindAddress
	}

	if c.BindPort <= 0 {
		return nil, ErrNoBindPort
	}

	cc, err := Consul()
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client: %w", err)
	}

	cs, err := connect.NewService(c.Service, cc)
	if err != nil {
		return nil, fmt.Errorf("failed to create connect service: %w", err)
	}

	return cs, nil
}
