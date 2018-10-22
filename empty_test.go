package main

import (
	"github.com/kami-zh/go-capturer"
	"os"
	"path/filepath"
	"testing"
)

func TestSearch_EmptyDirectories(t *testing.T) {
	t.Skip("Skip: Empty directories unable to register to Github")
}

func TestSearch_EmptyFiles(t *testing.T) {
	var exitCode int

	// update os.Args(dummy `empty.exe -f`)
	oldArgs := os.Args
	os.Args = []string{"empty.exe", "-f"}

	outStr := capturer.CaptureStdout(func() {
		exitCode = search()
	})

	if exitCode != 0 {
		t.Errorf("Unexpected exit code: %d", exitCode)
		os.Args = oldArgs
	}

	expectedStr := filepath.FromSlash("testdata/empty.txt\n") + filepath.FromSlash("testdata/test_dir1/empty.txt\n") + filepath.FromSlash("testdata/test_dir1/test_dir1_1/empty.txt\n")

	if outStr != expectedStr {
		t.Errorf("Unexpected output: %s", outStr)
		t.Logf("expected output: %s", expectedStr)
		os.Args = oldArgs
	}

	os.Args = oldArgs
}
