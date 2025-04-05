package agent

import (
	"context"
	"io"
	"log/slog"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/domain/hgraber"
	"github.com/gbh007/hgraber-next-agent-core/entities"
)

type loader interface {
	HasParser(ctx context.Context, u string) (bool, error)
	Load(ctx context.Context, URL string) (hgraber.BookParser, error)
	LoadImage(ctx context.Context, u string, bookUrl string) (io.ReadCloser, error)
	AllBooks(ctx context.Context, u string) ([]string, error)
	HProxyList(ctx context.Context, u url.URL) ([]entities.HProxyListUnit, error)
}

type UseCase struct {
	logger *slog.Logger

	loader loader
}

func New(
	logger *slog.Logger,
	loader loader,
) *UseCase {
	return &UseCase{
		logger: logger,
		loader: loader,
	}
}
