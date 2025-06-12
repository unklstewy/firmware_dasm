package analysis

import (
	"testing"

	fwtesting "github.com/unklstewy/firmware_dasm/pkg/analysis/testing"
)

func TestIdentifyFirmwareType(t *testing.T) {
	tests := []struct {
		name     string
		header   []byte
		expected FirmwareType
	}{
		{
			name:     "Baofeng firmware",
			header:   []byte{0x42, 0x46, 0x2D, 0x36, 0x31}, // "BF-61"
			expected: FirmwareTypeBaofeng,
		},
		{
			name:     "TYT firmware",
			header:   []byte{0x54, 0x59, 0x54, 0x33, 0x38}, // "TYT38"
			expected: FirmwareTypeTYT,
		},
		{
			name:     "Unknown firmware",
			header:   []byte{0x00, 0x01, 0x02, 0x03, 0x04},
			expected: FirmwareTypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IdentifyFirmwareType(tt.header)

			if result != tt.expected {
				t.Errorf("IdentifyFirmwareType() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractFirmwareSections(t *testing.T) {
	// This test is a placeholder - we need actual firmware samples
	t.Skip("Requires real firmware samples")

	analyzer := NewFirmwareAnalyzer()

	// Sample firmware data
	sampleData := fwtesting.LoadTestFirmware(t, "samples/firmware_sample.bin")

	sections, err := analyzer.ExtractSections(sampleData)
	if err != nil {
		t.Fatalf("ExtractSections failed: %v", err)
	}

	if len(sections) == 0 {
		t.Error("Expected to extract firmware sections, got none")
	}
}
