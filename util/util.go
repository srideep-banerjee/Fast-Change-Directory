package util

import (
	"path/filepath"
	"runtime"
)

func GetBinaryPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}