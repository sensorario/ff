package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// mockFileInfo implements os.FileInfo.  We're only interested in
// IsDir(), really.
type mockFileInfo struct {
	isDir bool
}

func (mockFileInfo) Name() string { return "" }

func (mockFileInfo) Size() int64 { return 0 }

func (mockFileInfo) Mode() os.FileMode { return 0 }

func (mockFileInfo) ModTime() time.Time { return time.Now() }

func (fi *mockFileInfo) IsDir() bool { return fi.isDir }

func (mockFileInfo) Sys() interface{} { return nil }

func TestFindRepository(t *testing.T) {

	t.Run("not found", func(t *testing.T) {
		repo, err := findRepositoryWithStatFn("/usr/local/src", func(string) (os.FileInfo, error) {
			return nil, os.ErrNotExist
		})
		assert.Equal(t, "", repo)
		assert.Equal(t, os.ErrNotExist, err)
	})

	t.Run("found", func(t *testing.T) {
		repo, err := findRepositoryWithStatFn("/home/user/src/project/sub", func(p string) (os.FileInfo, error) {
			if p == "/home/user/src/project/.git" {
				return &mockFileInfo{isDir: true}, nil
			}
			return nil, os.ErrNotExist
		})
		assert.Equal(t, "/home/user/src/project", repo)
		assert.Nil(t, err)
	})

	t.Run("found, not a dir", func(t *testing.T) {
		repo, err := findRepositoryWithStatFn("/home/user/src/project/sub", func(p string) (os.FileInfo, error) {
			if p == "/home/user/src/project/.git" {
				return &mockFileInfo{isDir: false}, nil
			}
			return nil, os.ErrNotExist
		})
		assert.Equal(t, "", repo)
		assert.NotNil(t, err)
	})

	t.Run("filesystem error", func(t *testing.T) {
		fsErr := fmt.Errorf("some other error")
		repo, err := findRepositoryWithStatFn("/home/user/src/project/sub", func(p string) (os.FileInfo, error) {
			return nil, fsErr
		})
		assert.Equal(t, "", repo)
		assert.Equal(t, fsErr, err)
	})

}
