package api

import (
	"context"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIHproxyParseBookPost(ctx context.Context, req *agentapi.APIHproxyParseBookPostReq) (agentapi.APIHproxyParseBookPostRes, error) {
	if c.parsingUseCases == nil {
		return &agentapi.APIHproxyParseBookPostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	details, err := c.parsingUseCases.HProxyBook(ctx, req.URL)
	if err != nil {
		return &agentapi.APIHproxyParseBookPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIHproxyParseBookPostOK{
		Name:       details.Name,
		URL:        details.URL,
		PreviewURL: OptURL(details.PreviewURL),
		PageCount:  details.PageCount,
		Pages: pkg.Map(details.Pages, func(p entities.HProxyBookDetailsPagesItem) agentapi.APIHproxyParseBookPostOKPagesItem {
			return agentapi.APIHproxyParseBookPostOKPagesItem{
				PageNumber: p.PageNumber,
				URL:        p.URL,
				Filename:   p.Filename,
			}
		}),
		Attributes: pkg.Map(details.Attributes, func(attr entities.HProxyBookDetailsAttributesItem) agentapi.APIHproxyParseBookPostOKAttributesItem {
			return agentapi.APIHproxyParseBookPostOKAttributesItem{
				Code: attr.Code,
				Values: pkg.Map(attr.Values, func(v entities.HProxyBookDetailsAttributesValueItem) agentapi.APIHproxyParseBookPostOKAttributesItemValuesItem {
					return agentapi.APIHproxyParseBookPostOKAttributesItemValuesItem{
						Name: v.Value,
						URL:  OptURL(v.URL),
					}
				}),
			}
		}),
	}, nil
}
