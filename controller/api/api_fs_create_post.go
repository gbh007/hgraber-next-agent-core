package api

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsCreatePost(ctx context.Context, req agentapi.APIFsCreatePostReq, params agentapi.APIFsCreatePostParams) (agentapi.APIFsCreatePostRes, error) {
	if c.fileUseCase == nil {
		return &agentapi.APIFsCreatePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	err := c.fileUseCase.Create(ctx, params.FileID, req.Data)
	if errors.Is(err, entities.FileAlreadyExistsError) {
		return &agentapi.APIFsCreatePostConflict{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentapi.APIFsCreatePostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIFsCreatePostNoContent{}, nil
}
