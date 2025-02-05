package api

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIExportArchivePost(ctx context.Context, req agentapi.APIExportArchivePostReq, params agentapi.APIExportArchivePostParams) (agentapi.APIExportArchivePostRes, error) {
	if c.exportUseCase == nil {
		return &agentapi.APIExportArchivePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	var u *url.URL

	if params.BookURL.IsSet() {
		u = &params.BookURL.Value
	}

	err := c.exportUseCase.Create(ctx, entities.ExportData{
		BookID:   params.BookID,
		BookName: params.BookName,
		Body:     req.Data,
		BookURL:  u,
	})
	if err != nil {
		return &agentapi.APIExportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIExportArchivePostNoContent{}, nil
}
