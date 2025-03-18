package lib

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DiskRepository implements Repository using the local filesystem.
type DiskRepository struct {
	BaseDir string
}

func NewDiskRepository(baseDir string) Repository {
	return &DiskRepository{BaseDir: baseDir}
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

	err = copyAndValidateSHA256Hash(tempFile, r, hash)
	if err != nil {
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
		if os.IsNotExist(err) {
			return nil, 0, ErrNotFound
		}
		return nil, 0, err
	}

	info, err := f.Stat()
	if err != nil {
		f.Close()
		if os.IsNotExist(err) {
			return nil, 0, ErrNotFound
		}
		return nil, 0, err
	}

	return f, info.Size(), nil
}

func copyAndValidateSHA256Hash(dst io.Writer, src io.Reader, expected string) error {
	hasher := sha256.New()
	mw := io.MultiWriter(dst, hasher)
	if _, err := io.Copy(mw, src); err != nil {
		return err
	}

	actual := fmt.Sprintf("%x", hasher.Sum(nil))
	if actual != expected {
		return &HashMismatchError{Expected: expected, Actual: actual}
	}

	return nil
}
