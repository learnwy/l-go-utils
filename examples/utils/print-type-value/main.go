package main

import (
	"fmt"
	"github.com/learnwy/l-go-utils/utils"
)

func main() {
	// Example usage of PrintTypeValue
	boolValue := true
	fmt.Printf("PrintTypeValue output for bool: %s\n", utils.PrintTypeValue(boolValue))

	intValue := 42
	fmt.Printf("PrintTypeValue output for int: %s\n", utils.PrintTypeValue(intValue))

	strValue := "Hello"
	fmt.Printf("PrintTypeValue output for string: %s\n", utils.PrintTypeValue(strValue))
}
