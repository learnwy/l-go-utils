// Package utils provides utility functions for Go development
package utils

import (
	"fmt"
)

// PrintTypeValue prints the value using %v format specifier
// It provides the default format for the value
func PrintTypeValue(v interface{}) string {
	return fmt.Sprintf("%v", v)
}