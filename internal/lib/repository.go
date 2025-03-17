package lib

import (
	"context"
	"io"
	"os"
	"path/filepath"
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
	// It must ensure atomicity: the content is either fully written (and made visible)
	// or not written at all. No partial writes should be observable.
	Put(ctx context.Context, hash string, r io.Reader) error

	// Get retrieves a ReadSeeker and the size (in bytes) for the content associated with the given hash.
	Get(ctx context.Context, hash string) (io.ReadSeeker, int64, error)
}

// DiskRepository implements Repository using the local filesystem.
type DiskRepository struct {
	BaseDir string
}

func NewDiskRepository(baseDir string) Repository {
	return &DiskRepository{BaseDir: baseDir}
}

func (d *DiskRepository) Exists(ctx context.Context, key string) bool {
	path := filepath.Join(d.BaseDir, key)
	_, err := os.Stat(path)
	return err == nil
}

// Put writes the content from r to a temporary file within BaseDir, flushes it to stable storage,
// and then atomically renames it to the final path (BaseDir/hash).
//
// The atomic rename guarantees that concurrent readers will either see the fully written file
// or no file at all—thereby ensuring no dirty/partial writes are exposed. Note that while
// File.Sync() flushes data to disk (equivalent to fsync on POSIX or FlushFileBuffers on Windows),
// the true durability of the write depends on the underlying OS and hardware configuration
func (d *DiskRepository) Put(ctx context.Context, hash string, r io.Reader) error {
	finalPath := filepath.Join(d.BaseDir, hash)

	tempFile, err := os.CreateTemp(d.BaseDir, "tmp-*")
	if err != nil {
		return err
	}

	tempPath := tempFile.Name()

	defer func() {
		tempFile.Close()
		os.Remove(tempPath)
	}()

	if _, err := io.Copy(tempFile, r); err != nil {
		return err
	}

	// Commit file contents to stable storage.
	if err := tempFile.Sync(); err != nil {
		return err
	}

	if err := tempFile.Close(); err != nil {
		return err
	}

	// Atomically move the temporary file to its final destination.
	return os.Rename(tempPath, finalPath)
}

func (d *DiskRepository) Get(ctx context.Context, hash string) (io.ReadSeeker, int64, error) {
	path := filepath.Join(d.BaseDir, hash)

	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, 0, err
	}

	return f, info.Size(), nil
}
