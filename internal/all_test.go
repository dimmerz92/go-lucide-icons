package internal

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetNewFiles(t *testing.T) {
	src := "./tests/src"
	htmlTarget := "./tests/html_target"
	templTarget := "./tests/templ_target"
	newFiles := []string{"tests/src/newfile1.svg", "tests/src/newfile2.svg"}

	output, err := getNewFiles(src, htmlTarget, templTarget)
	if err != nil {
		t.Fatalf("failed to get new files: %v", err)
	}

	if len(newFiles) != len(output["html"]) {
		t.Fatalf("getNewFiles returned %v\twant %v", output["html"], newFiles)
	}
	if len(newFiles) != len(output["templ"]) {
		t.Fatalf("getNewFiles returned %v\twant %v", output["templ"], newFiles)
	}

	for _, fileType := range output {
		for i := range fileType {
			if fileType[i] != newFiles[i] {
				t.Fatalf("%s should not be in returned array", fileType[i])
			}
		}
	}
}

func TestKebabtoPascalCase(t *testing.T) {
	if name := kebabToPascalCase("hello-world"); name != "HelloWorld" {
		t.Fatalf("expected HelloWorld, returned %s", name)
	}
	if name := kebabToPascalCase("hello"); name != "Hello" {
		t.Fatalf("expected Hello, returned %s", name)
	}
	if name := kebabToPascalCase("hello-world-long"); name != "HelloWorldLong" {
		t.Fatalf("expected HelloWorld, returned %s", name)
	}
	if name := kebabToPascalCase("hello--world"); name != "HelloWorld" {
		t.Fatalf("expected HelloWorld, returned %s", name)
	}
}

func TestGenerateTemplIcon(t *testing.T) {
	file := "./tests/src/newfile1.svg"
	target := "tests/templ_target"
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	if err := generateTemplIcon(file, target); err != nil {
		t.Fatalf("failed to generate templ icon: %v", err)
	}

	_, err = os.Stat(filepath.Join(cwd, target, "newfile1.templ"))
	if err != nil {
		t.Fatalf("templ icon did not get created: %v", err)
	}

	newFile, err := os.ReadFile(filepath.Join(target, "newfile1.templ"))
	if err != nil {
		t.Fatalf("failed to read newly generated templ file: %v", err)
	}

	t.Logf("inspect generated file contents:\n%s", string(newFile))

	err = os.Remove(filepath.Join(cwd, target, "newfile1.templ"))
	if err != nil {
		t.Logf("failed to remove the test generated templ file: %v", err)
	}
}

func TestGenerateHtmlIcon(t *testing.T) {
	file := "./tests/src/newfile1.svg"
	target := "tests/html_target"
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	if err := generateHtmlIcon(file, target); err != nil {
		t.Fatalf("failed to generate templ icon: %v", err)
	}

	_, err = os.Stat(filepath.Join(cwd, target, "newfile1.html"))
	if err != nil {
		t.Fatalf("html icon did not get created: %v", err)
	}

	newFile, err := os.ReadFile(filepath.Join(target, "newfile1.html"))
	if err != nil {
		t.Fatalf("failed to read newly generated html file: %v", err)
	}

	t.Logf("inspect generated file contents:\n%s", string(newFile))

	err = os.Remove(filepath.Join(cwd, target, "newfile1.html"))
	if err != nil {
		t.Logf("failed to remove the test generated html file: %v", err)
	}
}
