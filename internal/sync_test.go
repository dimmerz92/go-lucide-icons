package internal

import "testing"

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
