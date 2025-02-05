package api

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIParsingBookMultiPost(ctx context.Context, req *agentapi.APIParsingBookMultiPostReq) (agentapi.APIParsingBookMultiPostRes, error) {
	if c.parsingUseCases == nil {
		return &agentapi.APIParsingBookMultiPostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	result, err := c.parsingUseCases.MultiHandle(ctx, req.URL)
	if err != nil {
		return &agentapi.APIParsingBookMultiPostInternalServerError{
			InnerCode: ParseUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.BooksCheckResult{
		Result: convertBooksCheckResultResult(result),
	}, nil
}
