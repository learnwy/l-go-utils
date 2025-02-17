package utils

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintTypeValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"integer", 42, "(int, 42)\n"},
		{"string", "hello", "(string, hello)\n"},
		{"bool", true, "(bool, true)\n"},
		{"float", 3.14, "(float64, 3.14)\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call PrintTypeValue
			PrintTypeValue(tt.input)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read captured output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			if got != tt.expected {
				t.Errorf("PrintTypeValue() = %q, want %q", got, tt.expected)
			}
		})
	}
}
