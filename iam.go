package iam

import (
	"context"
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
func New() *IAM {
	return NewWithConfig(
		Config{
			EnableLogging: false,
		},
	)
}

// NewWithConfig creates an instance of IAM.
func NewWithConfig(config Config) *IAM {
	iam := &IAM{
		echo: echo.New(),
	}

	if config.EnableLogging {
		iam.echo.Use(middleware.Logger())
	}
	iam.echo.Use(middleware.Recover())
	iam.echo.Use(middleware.BodyFlush())

	return iam
}

// NewContext returns a Context instance.
func (iam *IAM) NewContext(r *http.Request, w http.ResponseWriter) echo.Context {
	return iam.echo.NewContext(r, w)
}

// ServeHTTP implements `http.Handler` interface, which serves HTTP requests.
func (iam *IAM) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	iam.echo.ServeHTTP(w, r)
}

// DefaultHTTPErrorHandler is the default HTTP error handler. It sends a JSON response
// with status code.
//
// NOTE: In case errors happens in middleware call-chain that is returning from handler (which did not return an error).
// When handler has already sent response (ala c.JSON()) and there is error in middleware that is returning from
// handler. Then the error that global error handler received will be ignored because we have already "commited" the
// response and status code header has been sent to the client.
func (iam *IAM) DefaultHTTPErrorHandler(err error, c echo.Context) {
	iam.echo.DefaultHTTPErrorHandler(err, c)
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

// Shutdown stops the server gracefully.
// It internally calls `http.Server#Shutdown()`.
func (iam *IAM) Shutdown(ctx context.Context) error {
	return iam.echo.Shutdown(ctx)
}
