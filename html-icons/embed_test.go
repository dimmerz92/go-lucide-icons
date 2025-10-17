package htmlicons_test

import (
	"testing"

	htmlicons "github.com/dimmerz92/go-lucide-icons/html-icons"
)

func TestGetHtmlFile(t *testing.T) {
	name := "a-arrow-up"
	data, err := htmlicons.GetHtmlFile(name)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal("returned empty data")
	}
}
