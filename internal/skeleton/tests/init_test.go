package tests

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/anonx/sunplate/internal/skeleton/assets/views"
	"github.com/anonx/sunplate/internal/skeleton/routes"

	"github.com/anonx/sunplate/controllers/templates"
	"github.com/anonx/sunplate/log"
)

var (
	s *httptest.Server
	h *http.Handler
)

func init() {
	// Server MUST be started from the root directory
	// of the project for correct loading of assets.
	os.Chdir("../")

	// Initialize a list of templates to use.
	templates.SetTemplatePaths(views.Root, views.List)

	// Build a handler and prepare a test server.
	h, err := routes.List.Build()
	log.AssertNil(err)
	s = httptest.NewUnstartedServer(h)
}
