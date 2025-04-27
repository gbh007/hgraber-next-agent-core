package api

import (
	"context"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIHproxyParseListPost(ctx context.Context, req *agentapi.APIHproxyParseListPostReq) (agentapi.APIHproxyParseListPostRes, error) {
	if c.parsingUseCases == nil {
		return &agentapi.APIHproxyParseListPostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	list, err := c.parsingUseCases.HProxyList(ctx, req.URL)
	if err != nil {
		return &agentapi.APIHproxyParseListPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIHproxyParseListPostOK{
		Results: pkg.Map(list, func(u entities.HProxyListUnit) agentapi.APIHproxyParseListPostOKResultsItem {
			var t agentapi.APIHproxyParseListPostOKResultsItemType

			switch u.Type {
			case entities.DetailsHProxyListUnitType:
				t = agentapi.APIHproxyParseListPostOKResultsItemTypeDetails
			case entities.ListHProxyListUnitType:
				t = agentapi.APIHproxyParseListPostOKResultsItemTypeList
			}

			return agentapi.APIHproxyParseListPostOKResultsItem{
				LinkURL:    u.LinkURL,
				Name:       OptString(u.Name),
				PreviewURL: OptURL(u.PreviewURL),
				Type:       t,
			}
		}),
	}, nil
}
