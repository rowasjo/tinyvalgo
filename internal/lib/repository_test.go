package lib

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/matryer/is"
)

const (
	example1_key   = "the-key"
	example1_value = `I am a little blob.`
	example1_size  = 19
)

func TestGetBlobInvalidHashReturns400(t *testing.T) {
	is := is.New(t)
	repo := NewDiskRepository(t.TempDir())

	_, _, err := repo.Get(context.Background(), "unknown key")
	is.True(err != nil)
}

func TestPutAndGetBlob(t *testing.T) {
	is := is.New(t)
	repo := NewDiskRepository(t.TempDir())
	ctx := context.Background()

	err := repo.Put(ctx, example1_key, strings.NewReader(example1_value))
	is.NoErr(err)

	reader, size, err := repo.Get(ctx, example1_key)
	is.NoErr(err)

	is.Equal(int64(19), size)

	data, err := io.ReadAll(reader)
	is.NoErr(err)

	is.Equal(example1_value, string(data))
}
