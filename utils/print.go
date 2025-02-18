// Package utils provides utility functions for Go development
package utils

import "fmt"

// FormatTypeValue returns a string representation of a value's type and content
// in the format "(type, value)". It accepts any type through the empty interface.
func FormatTypeValue(v interface{}) string {
	return fmt.Sprintf("(%T, %+v)", v, v)
}

// PrintTypeValue prints the type and value of the given interface to standard output
// in the format "(type, value)". It accepts any type through the empty interface.
func PrintTypeValue(v interface{}) {
	fmt.Printf("%s\n", FormatTypeValue(v))
}

// FormatConvertInfo creates a string describing the conversion between two values,
// showing their types and values. It accepts any two types through empty interfaces.
func FormatConvertInfo(from, to interface{}) string {
	return fmt.Sprintf("Convert %s to %s", FormatTypeValue(from), FormatTypeValue(to))
}

// PrintConvertInfo prints information about a type conversion to standard output,
// showing the original and converted values with their types.
func PrintConvertInfo(from, to interface{}) {
	fmt.Printf("%s\n", FormatConvertInfo(from, to))
}

// FormatSlice creates a string representation of a slice, showing its length, capacity,
// and elements. It accepts any type through the empty interface.
func FormatSlice(s []interface{}) string {
	return fmt.Sprintf("len=%d, cap=%d, %v", len(s), cap(s), s)
}

// PrintSlice prints the length, capacity, and elements of a slice to standard output.
// It accepts any type through the empty interface.
func PrintSlice(s []interface{}) {
	fmt.Printf("%s\n", FormatSlice(s))
}
