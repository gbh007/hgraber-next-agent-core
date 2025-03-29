package request

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const defaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"

type cache interface {
	SetString(ctx context.Context, u, v string)
	GetString(ctx context.Context, u string) (string, bool)
}

type Requester struct {
	logger *slog.Logger
	cache  cache
	client *http.Client
}

func New(
	logger *slog.Logger,
	timeout time.Duration,
	transport http.RoundTripper,
	cache cache,
) *Requester {
	if transport == nil {
		transport = http.DefaultTransport
	}

	return &Requester{
		client: &http.Client{
			Timeout: timeout,
			Transport: otelhttp.NewTransport(
				transport,
				otelhttp.WithPropagators(noopPropagator{}),
			),
		},
		logger: logger,
		cache:  cache,
	}
}

// RequestString запрашивает данные по урле и возвращает их строкой
func (r *Requester) RequestString(ctx context.Context, URL string) (string, error) {
	if r.cache != nil {
		v, ok := r.cache.GetString(ctx, URL)
		if ok {
			return v, nil
		}
	}

	buff, _, err := r.requestBuffer(ctx, URL, nil, nil)
	if err != nil {
		return "", err
	}

	if r.cache != nil {
		r.cache.SetString(ctx, URL, buff.String())
	}

	return buff.String(), nil
}

func (r *Requester) RequestStringWithRedirect(ctx context.Context, URL string) (string, string, error) {
	if r.cache != nil {
		v, ok := r.cache.GetString(ctx, URL)
		if ok {
			return v, URL, nil
		}
	}

	buff, resultURL, err := r.requestBuffer(ctx, URL, nil, nil)
	if err != nil {
		return "", "", err
	}

	if r.cache != nil {
		r.cache.SetString(ctx, resultURL, buff.String())
	}

	return buff.String(), resultURL, nil
}

// RequestBytes запрашивает данные по урле и возвращает их массивом байт
func (r *Requester) RequestBytes(ctx context.Context, URL string) ([]byte, error) {
	buff, _, err := r.requestBuffer(ctx, URL, nil, nil)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func (r *Requester) Request(ctx context.Context, URL string, headers http.Header) (io.ReadCloser, error) {
	// FIXME: работать с потоком напрямую
	buff, _, err := r.requestBuffer(ctx, URL, headers, nil)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(buff), nil
}

func (r *Requester) RequestPost(ctx context.Context, u string, headers http.Header, body io.Reader) ([]byte, error) {
	buff, _, err := r.requestBuffer(ctx, u, headers, body)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
