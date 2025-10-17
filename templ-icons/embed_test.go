package templicons_test

import (
	"testing"

	templicons "github.com/dimmerz92/go-lucide-icons/templ-icons"
)

func TestGetTemplFile(t *testing.T) {
	name := "a-arrow-up"
	data, err := templicons.GetTemplFile(name)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) == 0 {
		t.Fatal("returned empty data")
	}
}
