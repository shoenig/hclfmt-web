package config

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrNoBindAddress = errors.New("no bind address")
	ErrPortRange     = errors.New("port not within range")
)

const (
	defaultReadTimeout  = 3 * time.Second
	defaultWriteTimeout = 3 * time.Second
)

type Configuration struct {
	BindAddress   string        `hcl:"bind_address"`
	Port          int           `hcl:"port"`
	ReadTimeout   time.Duration `hcl:"read_timeout"`
	WriteTimeout  time.Duration `hcl:"write_timeout"`
	MaxRequestLen int64         `hcl:"max_request_length"`
}

func (c Configuration) Address() string {
	return fmt.Sprintf("%s:%d", c.BindAddress, c.Port)
}

func (c Configuration) Server(mux http.Handler) (*http.Server, error) {
	if c.BindAddress == "" {
		return nil, ErrNoBindAddress
	}

	if c.Port <= 1024 {
		return nil, ErrPortRange
	}

	readTimeout := c.ReadTimeout
	if readTimeout <= 0 {
		readTimeout = defaultReadTimeout
	}

	writeTimeout := c.WriteTimeout
	if writeTimeout <= 0 {
		writeTimeout = defaultWriteTimeout
	}

	server := &http.Server{
		Addr:         c.Address(),
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server, nil
}
