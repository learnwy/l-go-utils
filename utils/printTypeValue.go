// Package utils provides utility functions for Go development
package utils

import "fmt"

func FormatTypeValue(v interface{}) string {
	return fmt.Sprintf("(%T, %v)", v, v)
}

func PrintTypeValue(v interface{}) {
	fmt.Printf("%s\n", FormatTypeValue(v))
}
