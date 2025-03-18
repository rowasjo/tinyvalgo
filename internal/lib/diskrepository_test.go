package lib

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/matryer/is"
)

const (
	example1_key   = "bfb272e79d30466cf1af7c16739659e8b4e9b85b5075bdb922806c55035497cf"
	example1_value = "I am a little blob."
	example1_size  = 19
)

func TestGetBlobMissingHashReturns400(t *testing.T) {
	is := is.New(t)
	repo := NewDiskRepository(t.TempDir())

	_, _, err := repo.Get(context.Background(), "unknown key")
	is.Equal(err, ErrNotFound)
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
