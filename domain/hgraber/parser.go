package hgraber

import (
	"context"
	"errors"
	"io"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/entities"
)

var (
	InvalidLinkError      = errors.New("invalid link")
	UnknownAttributeError = errors.New("unknown attribute")
)

// Parser интерфейс для реализации парсеров для различных сайтов
type Parser interface {
	Name() string
	Load(ctx context.Context, u string) (BookParser, error)
	CanParse(u string) bool
	AllBooks(ctx context.Context, u string) ([]string, error)
	LoadImage(ctx context.Context, u string, bookUrl string) (io.ReadCloser, error)

	HProxyList(ctx context.Context, u url.URL) ([]entities.HProxyListUnit, error)
}

type BookParser interface {
	BookDetails(ctx context.Context, u url.URL) (entities.AgentBookDetails, error)
	HProxyBookDetails(ctx context.Context, u url.URL) (entities.HProxyBookDetails, error)
}
