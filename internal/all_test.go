package internal

import (
	"os"
	"testing"
)

func TestGetNewFiles(t *testing.T) {
	src := "./tests/src"
	target := "./tests/target"
	newFiles := []string{"tests/src/newfile1.svg", "tests/src/newfile2.svg"}

	files, err := getNewFiles(src, target)
	if err != nil {
		t.Fatalf("failed to get new files: %v", err)
	}

	if len(newFiles) != len(files) {
		t.Fatalf("getNewFiles returned %v\twant %v", files, newFiles)
	}

	for i := range files {
		if files[i] != newFiles[i] {
			t.Fatalf("%s should not be in returned array", files[i])
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
	target := "./tests/target/"
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	if err := generateTemplIcon(file, target); err != nil {
		t.Fatalf("failed to generate templ icon: %v", err)
	}

	if _, err := os.Stat(cwd + "/" + target + "newfile1.templ"); err != nil {
		t.Fatalf("templ icon did not get created: %v", err)
	}

	newFile, err := os.ReadFile(target + "newfile1.templ")
	if err != nil {
		t.Fatalf("failed to read newly generated templ file: %v", err)
	}

	t.Logf("inspect generated file contents:\n%s", string(newFile))

	if err := os.Remove(cwd + "/" + target + "newfile1.templ"); err != nil {
		t.Logf("failed to remove the test generated templ file: %v", err)
	}
}

func TestGenerateHtmlIcon(t *testing.T) {
	file := "./tests/src/newfile1.svg"
	target := "./tests/target/"
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}

	if err := generateHtmlIcon(file, target); err != nil {
		t.Fatalf("failed to generate templ icon: %v", err)
	}

	if _, err := os.Stat(cwd + "/" + target + "newfile1.html"); err != nil {
		t.Fatalf("html icon did not get created: %v", err)
	}

	newFile, err := os.ReadFile(target + "newfile1.html")
	if err != nil {
		t.Fatalf("failed to read newly generated html file: %v", err)
	}

	t.Logf("inspect generated file contents:\n%s", string(newFile))

	if err := os.Remove(cwd + "/" + target + "newfile1.html"); err != nil {
		t.Logf("failed to remove the test generated html file: %v", err)
	}
}
