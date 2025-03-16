package lib

import (
	"testing"

	"github.com/matryer/is"
)

func TestGetBlobInvalidHashReturns400(t *testing.T) {
	is := is.New(t)
	repo := NewDiskRepository(t.TempDir())

	_, _, err := repo.GetStream("unknown key")
	is.True(err != nil)
}
