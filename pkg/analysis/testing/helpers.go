package testing

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// GetTestDataPath returns the absolute path to the testdata directory
func GetTestDataPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "testdata")
}

// LoadTestFirmware loads a firmware file from the testdata directory
func LoadTestFirmware(t *testing.T, relativePath string) []byte {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), relativePath)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load test firmware %s: %v", path, err)
	}
	return data
}

// CompareResults compares analysis results with expected output
func CompareResults(t *testing.T, got, expected []byte) {
	t.Helper()

	if string(got) != string(expected) {
		t.Errorf("Analysis results don't match expected output")
	}
}
