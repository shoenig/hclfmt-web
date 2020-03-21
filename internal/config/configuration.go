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
	WebServer WebServer `hcl:"web_server"`
}

type WebServer struct {
	BindAddress   string        `hcl:"bind_address"`
	Port          int           `hcl:"port"`
	ReadTimeout   time.Duration `hcl:"read_timeout"`
	WriteTimeout  time.Duration `hcl:"write_timeout"`
	MaxRequestLen int64         `hcl:"max_request_length"`
}

func (ws WebServer) Address() string {
	return fmt.Sprintf("%s:%d", ws.BindAddress, ws.Port)
}

func (ws WebServer) Server(mux http.Handler) (*http.Server, error) {
	if ws.BindAddress == "" {
		return nil, ErrNoBindAddress
	}

	if ws.Port <= 1024 {
		return nil, ErrPortRange
	}

	readTimeout := ws.ReadTimeout
	if readTimeout <= 0 {
		readTimeout = defaultReadTimeout
	}

	writeTimeout := ws.WriteTimeout
	if writeTimeout <= 0 {
		writeTimeout = defaultWriteTimeout
	}

	server := &http.Server{
		Addr:         ws.Address(),
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server, nil
}
