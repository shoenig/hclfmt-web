package config

import (
	"errors"
	"fmt"
	"os"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/connect"
	"gophers.dev/pkgs/loggy"
)

var (
	ErrNoBindAddress = errors.New("no BIND address set")
	ErrNoBindPort    = errors.New("no PORT set")
	ErrNoService     = errors.New("no SERVICE set")
)

func logEnvironment(log loggy.Logger, name string) {
	value := os.Getenv(name)
	if value == "" {
		value = "<unset>"
	}
	log.Tracef("environment %s = %s", name, value)
}

func Consul(log loggy.Logger) (*consulapi.Client, error) {
	logEnvironment(log, "SERVICE")
	logEnvironment(log, "BIND")
	logEnvironment(log, "PORT")
	logEnvironment(log, "CONSUL_HTTP_ADDR")
	logEnvironment(log, "CONSUL_NAMESPACE")
	logEnvironment(log, "CONSUL_CACERT")
	logEnvironment(log, "CONSUL_CLIENT_CERT")
	logEnvironment(log, "CONSUL_CLIENT_KEY")
	logEnvironment(log, "CONSUL_HTTP_SSL")
	logEnvironment(log, "CONSUL_HTTP_SSL_VERIFY")
	logEnvironment(log, "CONSUL_TLS_SERVER_NAME")
	logEnvironment(log, "CONSUL_GRPC_ADDR")
	logEnvironment(log, "CONSUL_HTTP_TOKEN_FILE")
	consulConfig := consulapi.DefaultConfig()
	return consulapi.NewClient(consulConfig)
}

func Address() string {
	return fmt.Sprintf("%s:%s", os.Getenv("BIND"), os.Getenv("PORT"))
}

func GetService() (*connect.Service, error) {
	log := loggy.New("config")

	if bind := os.Getenv("BIND"); bind == "" {
		return nil, ErrNoBindAddress
	}

	if port := os.Getenv("PORT"); port == "" {
		return nil, ErrNoBindPort
	}

	if service := os.Getenv("SERVICE"); service == "" {
		return nil, ErrNoService
	}

	cc, err := Consul(log)
	if err != nil {
		return nil, fmt.Errorf("failed to create consul client: %w", err)
	}

	cs, err := connect.NewService(os.Getenv("SERVICE"), cc)
	if err != nil {
		return nil, fmt.Errorf("failed to create connect service: %w", err)
	}

	return cs, nil
}
