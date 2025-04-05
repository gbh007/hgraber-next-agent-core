package agent

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/pkg"
)

func (uc *UseCase) HProxyBook(ctx context.Context, u url.URL) (entities.HProxyBookDetails, error) {
	stringURL := u.String()

	parser, err := uc.loader.Load(ctx, stringURL)
	if err != nil {
		return entities.HProxyBookDetails{}, fmt.Errorf("load parser: %w", err)
	}

	details, err := parser.HProxyBookDetails(ctx, u)
	if err != nil {
		return entities.HProxyBookDetails{}, fmt.Errorf("parse: %w", err)
	}

	return details, nil
}

func (uc *UseCase) HProxyList(ctx context.Context, u url.URL) ([]entities.HProxyListUnit, error) {
	list, err := uc.loader.HProxyList(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("loader: %w", err)
	}

	list = pkg.SliceFilter(list, func(u entities.HProxyListUnit) bool {
		return u.Type != entities.UnknownHProxyListUnitType
	})

	return list, nil
}
