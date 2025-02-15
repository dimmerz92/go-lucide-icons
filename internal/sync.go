package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// getNewFiles returns a string slice of file names with relative path including
// `.svg` extension of new files in the src directory that are not in the target
// directory.
func getNewFiles(src, target string) ([]string, error) {
	targetFiles := make(map[string]struct{})
	newFiles := make([]string, 0)

	// generate a set from the target files
	targetDir, err := os.ReadDir(target)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s: %v", target, err)
	}

	for _, file := range targetDir {
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		targetFiles[name] = struct{}{}
	}

	// get the difference between the two directories
	srcDir, err := os.ReadDir(src)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %s: %v", target, err)
	}

	for _, file := range srcDir {
		if _, ok := targetFiles[strings.TrimSuffix(file.Name(), ".svg")]; !ok {
			newFiles = append(newFiles, filepath.Join(src, file.Name()))
		}
	}

	return newFiles, nil
}
