package funny

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// BacktracePathFromCurrentDir trace path in parents dir from current dir.
// Usage: trace some path like go.mod, .git, .gitignore...
func BacktracePathFromCurrentDir(path string) (string, os.FileInfo, error) {
	return BacktracePath(".", path)
}

// BacktracePath trace path in parents dir from given @from.
// Usage: trace some path like go.mod, .git, .gitignore...
func BacktracePath(from string, path string) (string, os.FileInfo, error) {
	path = strings.Trim(path, "/")

	from = TryGetAbsPath(from)
	flagRoot := false
	if TryGetAbsPath(from) == "/" {
		flagRoot = true
	}

	items, e := ioutil.ReadDir(from)
	if e != nil {
		return "", nil, e
	}

	for _, item := range items {
		if item.Name() == path {
			return TryGetAbsPath(filepath.Join(from, item.Name())), item, nil
		}
	}

	if flagRoot {
		return "", nil, os.ErrNotExist
	}

	return BacktracePath(filepath.Join(from, ".."), path)
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
