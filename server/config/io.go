package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// isFileExists reports whether the named file or directory exists.
func isFileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// OsAppRootPath returns default config path for the app.
func OsAppRootPath() string {
	name := AppServerName
	root := AppRootPathName
	dir, _ := os.UserConfigDir()

	switch runtime.GOOS {
	case "windows", "darwin":
		name = AppUsage
		root = strings.ToUpper(root)
	}

	return filepath.Join(dir, root, name)
}
