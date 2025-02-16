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
func getNewFiles(src, htmlTarget, templTarget string) (map[string][]string, error) {
	// generate a set from the target files
	htmlTargetDir, err := os.ReadDir(htmlTarget)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %s: %v", htmlTarget, err)
	}

	templTargetDir, err := os.ReadDir(templTarget)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %s: %v", templTarget, err)
	}

	targetFiles := make(map[string]struct{})
	targetDirs := append(htmlTargetDir, templTargetDir...)
	for _, file := range targetDirs {
		if ext := filepath.Ext(file.Name()); ext == ".html" || ext == ".templ" {
			targetFiles[file.Name()] = struct{}{}
		}
	}

	// get the difference between src and combined target directories
	srcDir, err := os.ReadDir(src)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %s: %v", src, err)
	}

	newFiles := make(map[string][]string)
	for _, file := range srcDir {
		// skip src file if not an svg
		if !strings.HasSuffix(file.Name(), ".svg") {
			continue
		}

		name := strings.TrimSuffix(file.Name(), ".svg")

		// if file is not in html files, add it
		if _, htmlFile := targetFiles[name+".html"]; !htmlFile {
			newFiles["html"] = append(
				newFiles["html"],
				filepath.Join(src, name+".svg"),
			)
		}

		// if file is not in templ files, add it
		if _, templFile := targetFiles[name+".templ"]; !templFile {
			newFiles["templ"] = append(
				newFiles["templ"],
				filepath.Join(src, name+".svg"),
			)
		}
	}

	return newFiles, nil
}

// kebabToPascalCase converts a kebab case string to a pascal case string.
func kebabToPascalCase(file string) string {
	var buf strings.Builder

	for _, sub := range strings.Split(file, "-") {
		if len(sub) < 1 {
			continue
		}
		buf.WriteString(strings.ToUpper(string(sub[0])))
		buf.WriteString(sub[1:])
	}

	return buf.String()
}
