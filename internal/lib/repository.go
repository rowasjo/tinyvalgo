package lib

import (
	"context"
	"fmt"
	"io"
)

// Interface for storing and retrieving file-like values using their SHA-256 hashes as keys.
// Implementations of this interface are expected to ensure that a stored value,
// identified by its hash, is immutable once stored. The Put method must ensure that
// a value is either fully written and consistent (via an atomic operation)
// or not visible at all.
type Repository interface {

	// Exists checks whether a value exists for the given hash.
	Exists(ctx context.Context, hash string) bool

	// Put stores the content read from r under the given hash.
	// It returns a *HashMismatchError if the computed hash doesn't match the provided hash,
	// indicating that the input data does not correspond to the expected hash.
	// The operation ensures atomicity: the content is either fully written (and made visible)
	// or not written at all. No partial writes should be observable.
	// Note that other errors (e.g. due to I/O or OS issues) may also be returned.
	Put(ctx context.Context, hash string, r io.Reader) error

	// Get retrieves a ReadSeeker and the size (in bytes) for the content associated with the given hash.
	Get(ctx context.Context, hash string) (io.ReadSeeker, int64, error)
}

type HashMismatchError struct {
	Expected string
	Actual   string
}

func (e *HashMismatchError) Error() string {
	return fmt.Sprintf("blob content hash mismatch: expected %s, got %s", e.Expected, e.Actual)
}
