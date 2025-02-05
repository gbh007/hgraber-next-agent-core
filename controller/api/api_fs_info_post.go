package api

import (
	"context"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Controller) APIFsInfoPost(ctx context.Context, req *agentapi.APIFsInfoPostReq) (agentapi.APIFsInfoPostRes, error) {
	if c.fileUseCase == nil {
		return &agentapi.APIFsInfoPostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	state, err := c.fileUseCase.State(ctx, req.IncludeFileIds.Value, req.IncludeFileSizes.Value)
	if err != nil {
		return &agentapi.APIFsInfoPostInternalServerError{
			InnerCode: FileUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIFsInfoPostOK{
		FileIds: state.FileIDs,
		TotalFileCount: agentapi.OptInt64{
			Value: state.TotalFileCount,
			Set:   state.TotalFileCount > 0,
		},
		TotalFileSize: agentapi.OptInt64{
			Value: state.TotalFileSize,
			Set:   state.TotalFileSize > 0,
		},
		AvailableSize: agentapi.OptInt64{
			Value: state.AvailableSize,
			Set:   state.AvailableSize > 0,
		},
		Files: pkg.Map(state.Files, func(raw entities.FSStateFile) agentapi.APIFsInfoPostOKFilesItem {
			return agentapi.APIFsInfoPostOKFilesItem{
				ID:        raw.ID,
				Size:      raw.Size,
				CreatedAt: raw.CreatedAt,
			}
		}),
	}, nil
}
