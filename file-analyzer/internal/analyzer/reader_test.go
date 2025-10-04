package analyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileAndReadAndCount(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	content := "hello world\nthis is Go\n" // words=5, line breaks=2

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	// ReadFile
	text, err := ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}
	if text != content {
		t.Fatalf("ReadFile content mismatch:\n got: %q\nwant: %q", text, content)
	}

	// ReadAndCount
	w, l, err := ReadAndCount(path)
	if err != nil {
		t.Fatalf("ReadAndCount error: %v", err)
	}
	if w != 5 || l != 2 {
		t.Fatalf("ReadAndCount = (%d,%d); want (5,2)", w, l)
	}
}

func TestReadFile_Err(t *testing.T) {
	_, err := ReadFile(filepath.Join(t.TempDir(), "nope.txt"))
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}