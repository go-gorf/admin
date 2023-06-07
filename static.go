package admin

import (
	"path/filepath"
	"runtime"
)

func StaticFilesPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), "frontend/out")
}
