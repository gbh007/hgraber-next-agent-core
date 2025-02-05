package api

import (
	"context"

	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIHighwayTokenCreatePost(ctx context.Context) (agentapi.APIHighwayTokenCreatePostRes, error) {
	if c.highwayUseCase == nil {
		return &agentapi.APIHighwayTokenCreatePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	token, vu, err := c.highwayUseCase.NewToken(ctx)
	if err != nil {
		return &agentapi.APIHighwayTokenCreatePostInternalServerError{
			InnerCode: HighwayUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIHighwayTokenCreatePostOK{
		ValidUntil: vu,
		Token:      token,
	}, nil
}
