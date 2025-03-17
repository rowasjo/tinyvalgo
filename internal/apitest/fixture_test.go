package apitest

import (
	"net/http"
	"testing"

	"github.com/rowasjo/tinyvalgo/internal/lib"
	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

func NewTestServer(t *testing.T) http.Handler {
	t.Helper()
	repo := lib.NewDiskRepository(t.TempDir())
	return tinyvalapi.NewServer(repo)
}
