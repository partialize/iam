package iam

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIAM(t *testing.T) {
	iam, err := New()

	assert.NoError(t, err)
	assert.NotNil(t, iam)
}

func TestIAM_Start(t *testing.T) {
	iam, err := New()
	errChan := make(chan error)

	go func() {
		err := iam.Start(":0")
		if err != nil {
			errChan <- err
		}
	}()

	err = waitForServerStart(iam, errChan)

	assert.NoError(t, err)
	assert.NoError(t, iam.Close())
}

func TestIAM_ServeHTTP(t *testing.T) {
	iam, err := New()
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	iam.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}

func waitForServerStart(iam *IAM, errChan <-chan error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			addr := iam.ListenerAddr()
			if addr != nil && strings.Contains(addr.String(), ":") {
				return nil // was started
			}
		case err := <-errChan:
			if err == http.ErrServerClosed {
				return nil
			}
			return err
		}
	}
}
