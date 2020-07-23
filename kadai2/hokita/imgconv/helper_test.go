package imgconv

import (
	"os"
	"testing"
)

func checkAndDeleteFile(t *testing.T, file string) {
	t.Helper()

	if _, err := os.Stat(file); err != nil {
		t.Errorf(`"%v" was not found`, file)
	}

	if err := os.Remove(file); err != nil {
		t.Fatal(err)
	}
}
