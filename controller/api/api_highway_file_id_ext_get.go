package api

import (
	"context"
	"errors"
	"mime"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIHighwayFileIDExtGet(ctx context.Context, params agentapi.APIHighwayFileIDExtGetParams) (agentapi.APIHighwayFileIDExtGetRes, error) {
	if c.highwayUseCase == nil {
		return &agentapi.APIHighwayFileIDExtGetBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	if params.Token == "" {
		return &agentapi.APIHighwayFileIDExtGetUnauthorized{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	err := c.highwayUseCase.ValidateToken(ctx, params.Token)
	if err != nil {
		return &agentapi.APIHighwayFileIDExtGetForbidden{
			InnerCode: HighwayUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	body, err := c.highwayUseCase.Get(ctx, params.ID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &agentapi.APIHighwayFileIDExtGetNotFound{
			InnerCode: HighwayUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentapi.APIHighwayFileIDExtGetInternalServerError{
			InnerCode: HighwayUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	// Это не самый правильный и ленивый костыль, но пока его будет достаточно
	contentType := mime.TypeByExtension("." + params.Ext)

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	return &agentapi.APIHighwayFileIDExtGetOKHeaders{
		ContentType: contentType,
		Response: agentapi.APIHighwayFileIDExtGetOK{
			Data: body,
		},
	}, nil
}
