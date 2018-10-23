package main

import (
	"github.com/kami-zh/go-capturer"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var emptyFiles = []string{
	filepath.FromSlash("testdata/empty.txt"),
	filepath.FromSlash("testdata/test_dir1/empty.txt"),
	filepath.FromSlash("testdata/test_dir1/test_dir1_1/empty.txt"),
}

func TestSearch_EmptyDirectories(t *testing.T) {
	t.Skip("Skip: Empty directories unable to register to Github")
}

func TestSearch_EmptyFiles(t *testing.T) {
	var exitCode int

	// update os.Args(dummy `empty -f`)
	oldArgs := os.Args
	os.Args = []string{"empty", "-f"}

	outStr := capturer.CaptureStdout(func() {
		exitCode = search()
	})

	if exitCode != 0 {
		t.Errorf("Unexpected exit code: %d", exitCode)
		os.Args = oldArgs
	}

	expectedStr := strings.Join(emptyFiles, "\n") + "\n"
	t.Logf("Expected output: %s", expectedStr)

	if outStr != expectedStr {
		t.Errorf("Unexpected output: %s", outStr)
	}

	os.Args = oldArgs
}

func TestSearch_EmptyFilesPrintAbsolutePath(t *testing.T) {
	var exitCode int

	// update os.Args(dummy `empty -f -a`)
	oldArgs := os.Args
	os.Args = []string{"empty", "-f", "-a"}

	outStr := capturer.CaptureStdout(func() {
		exitCode = search()
	})

	if exitCode != 0 {
		t.Errorf("Unexpected exit code: %d", exitCode)
		os.Args = oldArgs
	}

	var expectedStr string
	for _, path := range emptyFiles {
		absPath, err := filepath.Abs(path)
		if err != nil {
			t.Errorf("[ERROR]: %v\n", err)
		}
		expectedStr = expectedStr + absPath + "\n"
	}
	t.Logf("Expected output: %s", expectedStr)

	if outStr != expectedStr {
		t.Errorf("Unexpected output: %s", outStr)
	}

	os.Args = oldArgs
}
