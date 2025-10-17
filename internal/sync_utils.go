package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// FileSet returns a set of names of files of type 'ext' from the specified 'path' with 'ext' trimmed.
func FileSet(path, ext string) map[string]struct{} {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	icons := make(map[string]struct{})
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		entryExt := strings.ToLower(filepath.Ext(entry.Name()))
		if entryExt == strings.ToLower(ext) {
			name := strings.TrimSuffix(strings.ToLower(entry.Name()), entryExt)
			icons[name] = struct{}{}
		}
	}

	return icons
}

// DiffSet returns the difference set of `setA` minus `setB`.
func DiffFileSet(setA, setB map[string]struct{}) map[string]struct{} {
	diff := make(map[string]struct{})
	for file := range setA {
		if _, ok := setB[file]; !ok {
			diff[file] = struct{}{}
		}
	}

	return diff
}
