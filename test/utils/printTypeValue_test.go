package utils

import (
	"testing"

	"github.com/learnwy/l-go-utils/utils"
)

func TestPrintTypeValue(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"integer", 42, "42"},
		{"string", "hello", "hello"},
		{"bool", true, "true"},
		{"float", 3.14, "3.14"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.PrintTypeValue(tt.input); got != tt.expected {
				t.Errorf("PrintTypeValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}