package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const HTML = `{{ define "%s" }}
%s
  {{ range $value := . }}
    {{ $value }}
  {{ end }}
>
%s
{{ end }}`

// generateHtmlIcon generates a html template icon from the given funcString
// template string and saves the templ file at the given target directory.
func generateHtmlIcon(file, target string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read src file %s: %v", file, err)
	}

	svgParts := strings.SplitN(string(content), ">", 2)
	if len(svgParts) != 2 {
		return fmt.Errorf("malformed file: %v", svgParts)
	}

	fileName := strings.TrimSuffix(filepath.Base(file), ".svg")

	tpl := fmt.Sprintf(HTML, fileName, svgParts[0], svgParts[1])

	err = os.WriteFile(
		filepath.Join(target, fileName+".html"),
		[]byte(tpl),
		0644,
	)
	if err != nil {
		return fmt.Errorf("failed to generate html/template icon: %v", err)
	}

	return nil
}
