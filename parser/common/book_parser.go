package common

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/domain/hgraber"
	"github.com/gbh007/hgraber-next-agent-core/entities"
)

// Проверка соответствия базового типа
var (
	_ hgraber.BookParser = (*BookParser)(nil)
)

type BookParser struct{}

func (BookParser) BookDetails(ctx context.Context, u url.URL) (entities.AgentBookDetails, error) {
	return entities.AgentBookDetails{
		URL: u,
	}, nil
}

func (BookParser) HProxyBookDetails(ctx context.Context, u url.URL) (entities.HProxyBookDetails, error) {
	return entities.HProxyBookDetails{
		URL: u,
	}, nil
}
