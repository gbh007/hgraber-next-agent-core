package loader

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"time"

	"github.com/gbh007/hgraber-next-agent-core/domain/hgraber"
	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next-agent-core/metrics"
	"github.com/gbh007/hgraber-next-agent-core/parser/hgraber_local"
	"github.com/gbh007/hgraber-next-agent-core/parser/mock"
	"github.com/gbh007/hgraber-next-agent-core/request"
)

type Cache interface {
	SetString(ctx context.Context, u, v string)
	GetString(ctx context.Context, u string) (string, bool)
}

type Loader struct {
	parsers []hgraber.Parser
}

func NewDefaultParsers(
	logger *slog.Logger,
	hgToken string,
	timeout time.Duration,
	enabledParsers []string,
	cache Cache,
) []hgraber.Parser {
	requester := request.New(logger, timeout, nil, cache)

	parsers := make([]hgraber.Parser, 0, len(enabledParsers))

	for _, code := range enabledParsers {
		var p hgraber.Parser

		switch code {
		case "mock":
			p = mock.New(requester)

		case "hgraber_local":
			p = hgraber_local.New(requester, hgToken)

		default:
			logger.Warn(
				"unknown parser code",
				slog.String("code", code),
			)

			continue
		}

		parsers = append(parsers, p)
	}

	return parsers
}

func New(
	parsers []hgraber.Parser,
) *Loader {
	return &Loader{
		parsers: parsers,
	}
}

func (l *Loader) getParser(u string) (hgraber.Parser, error) {
	for _, p := range l.parsers {
		if p.CanParse(u) {
			return p, nil
		}
	}

	return nil, hgraber.InvalidLinkError
}

func (l *Loader) HasParser(ctx context.Context, u string) (bool, error) {
	_, err := l.getParser(u)
	if errors.Is(err, hgraber.InvalidLinkError) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("get parser: %w", err)
	}

	return true, nil
}

func (l *Loader) Load(ctx context.Context, u string) (hgraber.BookParser, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return nil, fmt.Errorf("get parser: %w", err)
	}

	defer func() {
		metrics.RegisterParserActionTime("loader_load", p.Name(), time.Since(startAt))
	}()

	bookParser, err := p.Load(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}

	return bookParser, nil
}

func (l *Loader) LoadImage(ctx context.Context, u string, bookUrl string) (io.ReadCloser, error) {
	startAt := time.Now()

	p, err := l.getParser(bookUrl)
	if err != nil {
		return nil, fmt.Errorf("get parser: %w", err)
	}

	defer func() {
		metrics.RegisterParserActionTime("loader_load_image", p.Name(), time.Since(startAt))
	}()

	data, err := p.LoadImage(ctx, u, bookUrl)
	if err != nil {
		return nil, fmt.Errorf("load image: %w", err)
	}

	return data, nil
}

func (l *Loader) AllBooks(ctx context.Context, u string) ([]string, error) {
	startAt := time.Now()

	p, err := l.getParser(u)
	if err != nil {
		return nil, err
	}

	defer func() {
		metrics.RegisterParserActionTime("loader_all_books", p.Name(), time.Since(startAt))
	}()

	data, err := p.AllBooks(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("load books: %w", err)
	}

	return data, nil
}

func (l *Loader) HProxyList(ctx context.Context, u url.URL) ([]entities.HProxyListUnit, error) {
	startAt := time.Now()

	p, err := l.getParser(u.String())
	if err != nil {
		return nil, err
	}

	defer func() {
		metrics.RegisterParserActionTime("loader_hproxy_list", p.Name(), time.Since(startAt))
	}()

	data, err := p.HProxyList(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("load hproxy list: %w", err)
	}

	return data, nil
}
