package egu

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// RuntimeCaller return code trace info
func RuntimeCaller() []string {
	trace := []string{}
	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok {
			trace = append(trace, fmt.Sprintf("[Runtime]%s[%d]: \n", file, line))
		}
	}

	return trace
}


// GetCurrentPath return compiled executable file absolute path
func GetCurrentPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// GetCurrentDir return compiled executable file directory
func GetCurrentDir() string {
	return filepath.Dir(GetCurrentPath())
}


// PathExist check the given path exists
func PathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Mkdir Create the DIRECTORY(ies), if they do not already exist
// parents no error if existing, make parent directories as needed
func Mkdir(dir string, parents bool) error {
	if PathExist(dir) {
		return nil
	}

	if parents {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return os.Mkdir(dir, os.ModePerm)
}


// SplitPathExt return filename and ext of path
func SplitPathExt(path string) (string, string) {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[:i], path[i:]
		}
	}
	return path, ""
}