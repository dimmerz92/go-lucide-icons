package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const TEMPL = `package icons

templ %s(props IconProps) {
    %s
		if props.ID != "" {
			id={ props.ID }
		}
		if props.Class != "" {
			class={ props.Class }
		}
		if props.Style != "" {
			style={ props.Style }
		}
		{ props.Attributes... }
    >
    %s
}`

// generateTemplIcon generates a templ template icon from the given funcString
// template string and saves the templ file at the given target directory.
func generateTemplIcon(file, target string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read src file %s: %v", file, err)
	}

	svgParts := strings.SplitN(string(content), ">", 2)
	if len(svgParts) != 2 {
		return fmt.Errorf("malformed file: %v", svgParts)
	}

	fileName := strings.TrimSuffix(filepath.Base(file), ".svg")
	funcName := kebabToPascalCase(fileName)

	templ := fmt.Sprintf(TEMPL, funcName, svgParts[0], svgParts[1])

	err = os.WriteFile(
		filepath.Join(target, fileName+".templ"),
		[]byte(templ),
		0644,
	)
	if err != nil {
		return fmt.Errorf("failed to generate templ icon: %v", err)
	}

	return nil
}
