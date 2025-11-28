package pkg

/**
* This File Helps find Valid .dll or .so plugins that match valid plugins
* in the cwd ./plugins folder.
* It is used by the manager application to load plugins dynamically.
 */

import (
	"os"
	"path/filepath"
	"strings"
)

// FindPlugins recursively looks for .dll files in the root directory
func FindPlugins(root string) ([]string, error) {
	var plugins []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Look for files ending in .dll
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".dll") {
			plugins = append(plugins, path)
		}
		return nil
	})

	return plugins, err
}
