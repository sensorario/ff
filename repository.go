package ff

import (
	"fmt"
	"os"
	"path"
)

// findRepository looks for a repository directory starting from the
// given directory, going up one level until the root directory or
// until the repository is found.  Returns os.NotExists in case the
// algorithm ran successfully but found no repository.
func FindRepository(dir string) (pathname string, err error) {
	return findRepositoryWithStatFn(dir, os.Stat)
}

func findRepositoryWithStatFn(dir string, statFn func(string) (os.FileInfo, error)) (pathname string, err error) {
	candidate := path.Join(dir, ".git")
	info, err := statFn(candidate)
	if err == nil {
		if info.IsDir() {
			return dir, nil
		}
		return "", fmt.Errorf("%q exists but is not a directory", candidate)
	}
	if os.IsNotExist(err) && dir != "/" {
		return findRepositoryWithStatFn(path.Clean(path.Join(dir, "..")), statFn)
	}
	return "", err
}
