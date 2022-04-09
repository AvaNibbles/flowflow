package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/avanibbles/flowflow/internal"
	httptransport "github.com/go-openapi/runtime/client"
	"go.uber.org/zap"

	"github.com/avanibbles/flowflow/pkg/client"
	"github.com/avanibbles/flowflow/pkg/config"
)

func NewApiClient(config *config.Config) (*client.Flowflowclient, error) {
	if config.Client == nil {
		return nil, errors.New("client not configured")
	}

	clientUrl, err := url.Parse(config.Client.Host)
	if err != nil {
		return nil, err
	}

	transport := httptransport.New(clientUrl.Host, clientUrl.Path, []string{clientUrl.Scheme})
	setMiddleware(transport)

	return client.New(transport, nil), nil
}

type apiClientMiddleware struct {
	inner     http.RoundTripper
	userAgent string
	debug     bool
	logger    *zap.Logger
}

func (u *apiClientMiddleware) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", u.userAgent)

	if u.debug {
		u.logger.Info("api client request",
			zap.String("path", r.URL.EscapedPath()),
			zap.String("method", r.Method),
			zap.String("host", r.Host))

		start := time.Now()
		resp, err := u.inner.RoundTrip(r)
		elapsed := time.Since(start)

		if err != nil {
			u.logger.Info("api client response error", zap.Error(err))
		} else {
			u.logger.Info("api client response",
				zap.String("status", resp.Status),
				zap.Duration("elapsed", elapsed))
		}

		return resp, err
	}

	return u.inner.RoundTrip(r)
}

func setMiddleware(r *httptransport.Runtime) {
	rt := &apiClientMiddleware{
		inner:     r.Transport,
		userAgent: fmt.Sprintf("flowflow-golang-client/%s", internal.Version),
		debug:     true,
		logger:    internal.GetLogger().With(zap.String("component", "apiClient")),
	}

	r.Transport = rt
}
