/*
testdata
│  empty.txt
│  not_empty.txt
│
├─.svn
│  │  empty.txt
│  │  not_empty.txt
│  │
│  ├─empty_dir (unable to register to Github)
│  └─not_empty_dir
│          empty.txt
│          not_empty.txt
│
├─empty_dir (unable to register to Github)
└─not_empty_dir
    │  empty.txt
    │  not_empty.txt
    │
    ├─empty_dir (unable to register to Github)
    └─not_empty_dir
            empty.txt
            not_empty.txt
*/
package main

import (
	"flag"
	"github.com/kami-zh/go-capturer"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var emptyFiles = []string{
	filepath.FromSlash("testdata/empty.txt"),
	filepath.FromSlash("testdata/not_empty_dir/empty.txt"),
	filepath.FromSlash("testdata/not_empty_dir/not_empty_dir/empty.txt"),
}

var excludeNotEmptyDirEmptyFiles = []string{
	filepath.FromSlash("testdata/.svn/empty.txt"),
	filepath.FromSlash("testdata/empty.txt"),
}

var (
	exitCode = 0
	oldArgs  = os.Args
)

func TestSearch_EmptyDirectories(t *testing.T) {
	t.Skip("Skip: Empty directories unable to register to Github")
}

func TestSearch_EmptyFiles(t *testing.T) {
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

	AfterEachTest()
}

func TestSearch_EmptyFilesPrintAbsolutePath(t *testing.T) {
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

	AfterEachTest()
}

func TestSearch_EmptyFilesIgnoreDirectory(t *testing.T) {
	// update os.Args(dummy `empty -f -i not_empty_dir`)
	oldArgs := os.Args
	os.Args = []string{"empty", "-f", "-i", "not_empty_dir"}

	outStr := capturer.CaptureStdout(func() {
		exitCode = search()
	})

	if exitCode != 0 {
		t.Errorf("Unexpected exit code: %d", exitCode)
		os.Args = oldArgs
	}

	expectedStr := strings.Join(excludeNotEmptyDirEmptyFiles, "\n") + "\n"
	t.Logf("Expected output: %s", expectedStr)

	if outStr != expectedStr {
		t.Errorf("Unexpected output: %s", outStr)
	}

	AfterEachTest()
}

func AfterEachTest() {
	// set flag default value
	os.Args = []string{"empty", "-f=0", "-a=0", "-i", "^(.git|.svn)$"}
	flag.Parse()
	os.Args = oldArgs
}
