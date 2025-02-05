package api

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsGetGet(ctx context.Context, params agentapi.APIFsGetGetParams) (agentapi.APIFsGetGetRes, error) {
	if c.fileUseCase == nil {
		return &agentapi.APIFsGetGetBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	body, err := c.fileUseCase.Get(ctx, params.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &agentapi.APIFsGetGetNotFound{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentapi.APIFsGetGetInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	// FIXME: работать с типом контента как в основном сервере
	return &agentapi.APIFsGetGetOK{
		Data: body,
	}, nil
}
