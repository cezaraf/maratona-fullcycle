package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HelloHandler - /hello
func HelloHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Hello Full Cycle"}))
}
