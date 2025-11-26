package config

import (
	"os"
	"path/filepath"
	"strings"
)

// Search files
func searchFile(dirPath string) []string {

	// Result search files
	var result []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip folders
		if info.IsDir() {
			return nil
		}

		// Same file extension
		if strings.ToLower(filepath.Ext(path)) == "."+strings.ToLower(dirPath) {
			result = append(result, path)
		}

		return nil
	})
	return result
}

// PNG
func SearchPngFile() []string {
	return searchFile("png")
}

// JPG
func SearchJpgFile() []string {
	return searchFile("jpg")
}
