package iam

import (
	"github.com/partialize/echo-slim/v4"
	"github.com/partialize/echo-slim/v4/middleware"
	"net"
	"net/http"
)

type (
	// Config defines the config for IAM.
	Config struct {
		EnableLogging bool
	}

	// IAM is the top-level framework instance.
	IAM struct {
		echo *echo.Echo
	}
)

// New creates an instance of IAM.
func New() (*IAM, error) {
	return NewWithConfig(
		Config{
			EnableLogging: false,
		},
	)
}

// NewWithConfig creates an instance of IAM.
func NewWithConfig(config Config) (*IAM, error) {
	e := echo.New()
	e.HideBanner = true

	if config.EnableLogging {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	return &IAM{
		echo: e,
	}, nil
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
func (iam *IAM) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iam.echo.ServeHTTP(w, r)
}

// ListenerAddr returns net.Addr for Listener
func (iam *IAM) ListenerAddr() net.Addr {
	return iam.echo.ListenerAddr()
}

// Start starts an HTTP server.
func (iam *IAM) Start(address string) error {
	return iam.echo.Start(address)
}

// Close immediately stops the server.
// It internally calls `http.Server#Close()`.
func (iam *IAM) Close() error {
	return iam.echo.Close()
}
