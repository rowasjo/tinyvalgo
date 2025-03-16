package lib

import (
	"io"
	"os"
	"path/filepath"
)

// Interface for storing and retrieving file-like values using strings as keys.
type Repository interface {

	// Check if value exists
	Exists(key string) bool

	// Sets value from stream
	SetFromStream(key string, r io.Reader) error

	// Returns an io.ReadSeeker and the size in bytes for the value associated with the key.
	// If the value doesn't exist or the operation fails, it returns an error.
	GetStream(key string) (io.ReadSeeker, int64, error)
}

// DiskRepository implements Repository using the local filesystem.
type DiskRepository struct {
	BaseDir string
}

func NewDiskRepository(baseDir string) Repository {
	return &DiskRepository{BaseDir: baseDir}
}

func (d *DiskRepository) Exists(key string) bool {
	path := filepath.Join(d.BaseDir, key)
	_, err := os.Stat(path)
	return err == nil
}

func (d *DiskRepository) SetFromStream(key string, r io.Reader) error {
	path := filepath.Join(d.BaseDir, key)

	// returns error if repository dir does not exist
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// TODO: atomic move on no error
	_, err = io.Copy(f, r)
	return err
}

func (d *DiskRepository) GetStream(key string) (io.ReadSeeker, int64, error) {
	path := filepath.Join(d.BaseDir, key)

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
