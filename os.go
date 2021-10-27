package funny

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// BacktracePath trace path from parents directories.
// Usage: trace some path like go.mod, .git, .gitignore...
func BacktracePath(path string) (string, os.FileInfo, error) {
	path = strings.Trim(path, "/")

	return backtracePath(".", path)
}

func backtracePath(dir string, path string) (string, os.FileInfo, error) {
	flagRoot := false
	if TryGetAbsPath(dir) == "/" {
		flagRoot = true
	}

	items, e := ioutil.ReadDir(dir)
	if e != nil {
		return "", nil, e
	}

	for _, item := range items {
		if item.Name() == path {
			return TryGetAbsPath(filepath.Join(dir, item.Name())), item, nil
		}
	}

	if flagRoot {
		return "", nil, e
	}

	return backtracePath(filepath.Dir(dir), path)
}

// TryGetAbsPath try to get absolute path of path but ignore the error.
// When got an error, it would return "" instead.
func TryGetAbsPath(path string) string {
	if path == "" {
		panic("got empty path param")
	}
	abs, e := filepath.Abs(path)
	if e != nil {
		return ""
	}

	return abs
}

//TryGetwd try to get current work directory and ignore the error.
// When got an error, it would return "" instead.
func TryGetwd() string {
	wd, e := os.Getwd()
	if e != nil {
		return ""
	}

	return wd
}
