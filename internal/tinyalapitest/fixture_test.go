package tinyvalapitest

import (
	"net/http"
	"testing"

	"github.com/rowasjo/tinyvalgo/internal/lib"
	"github.com/rowasjo/tinyvalgo/internal/tinyvalapi"
)

func NewTestApp(t *testing.T) http.Handler {
	t.Helper()
	repo := lib.NewDiskRepository(t.TempDir())
	return tinyvalapi.NewApp(repo)
}
