package api

import (
	"context"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/agentapi"
)

func (c *Controller) APIImportArchivePost(ctx context.Context, req agentapi.APIImportArchivePostReq, params agentapi.APIImportArchivePostParams) (agentapi.APIImportArchivePostRes, error) {
	if c.importUseCase == nil {
		return &agentapi.APIImportArchivePostBadRequest{
			InnerCode: ValidationCode,
			Details:   agentapi.NewOptString("unsupported api"),
		}, nil
	}

	var u *url.URL

	if params.BookURL.IsSet() {
		u = &params.BookURL.Value
	}

	err := c.importUseCase.Create(ctx, entities.ImportData{
		BookID:   params.BookID,
		BookName: params.BookName,
		Body:     req.Data,
		BookURL:  u,
	})
	if err != nil {
		return &agentapi.APIImportArchivePostInternalServerError{
			InnerCode: ExportUseCaseCode,
			Details:   agentapi.NewOptString(err.Error()),
		}, nil
	}

	return &agentapi.APIImportArchivePostNoContent{}, nil
}
