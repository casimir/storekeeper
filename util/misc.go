package util

import (
	"os"
	"runtime"
)

func ApplicationPath() string {
	switch runtime.GOOS {
	case "darwin":
		return os.Getenv("HOME") + "/Library/Application Support/storekeeper"
	case "linux":
		return os.Getenv("HOME") + "/.storekeeper"
	case "windows":
		return os.Getenv("APPDATA") + "/storekeeper"
	default:
		return os.Getenv("TMP") + "/storekeeper"
	}
}
