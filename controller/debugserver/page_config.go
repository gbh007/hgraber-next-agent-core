package debugserver

import (
	"bytes"
	"net/http"

	"github.com/gbh007/hgraber-next-agent-core/config"
	"github.com/labstack/echo/v4"
)

func (cnt *Controller[T]) pageConfig(c echo.Context) error {
	buff := &bytes.Buffer{}

	err := config.ExportToWriter(cnt.config, buff)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "config.html.gotmpl", buff.String())
}
