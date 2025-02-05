package api

import (
	"context"
	"errors"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIFsDeletePost(ctx context.Context, req *agentapi.APIFsDeletePostReq) (agentapi.APIFsDeletePostRes, error) {
	if c.fileUseCase == nil {
		return &agentapi.APIFsDeletePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	err := c.fileUseCase.Delete(ctx, req.FileID)
	if errors.Is(err, entities.FileNotFoundError) {
		return &agentapi.APIFsDeletePostNotFound{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	if err != nil {
		return &agentapi.APIFsDeletePostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIFsDeletePostNoContent{}, nil
}
