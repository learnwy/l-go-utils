package utils

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
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

func TestPrintSlice(t *testing.T) {
	// Test cases for interface{} slices
	tests := []struct {
		name     string
		input    []interface{}
		expected string
	}{{
		name:     "empty_slice",
		input:    []interface{}{},
		expected: "len=0, cap=0, []\n",
	}, {
		name:     "mixed_types",
		input:    []interface{}{1, "hello", true, 3.14},
		expected: "len=4, cap=4, [1 hello true 3.14]\n",
	}, {
		name:     "numbers_only",
		input:    []interface{}{1, 2, 3},
		expected: "len=3, cap=3, [1 2 3]\n",
	}, {
		name:     "strings_only",
		input:    []interface{}{"a", "b", "c"},
		expected: "len=3, cap=3, [a b c]\n",
	}, {
		name:     "with_nil",
		input:    []interface{}{nil, 1, "test"},
		expected: "len=3, cap=3, [<nil> 1 test]\n",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call PrintSlice
			PrintSlice(tt.input)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read captured output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			if got != tt.expected {
				t.Errorf("PrintSlice() = %q, want %q", got, tt.expected)
			}
		})
	}

	// Test cases for typed slices
	type Person struct {
		Name string
		Age  int
	}

	t.Run("int_slice", func(t *testing.T) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		PrintSlice([]int{1, 2, 3})

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(r)
		got := buf.String()
		expected := "len=3, cap=3, [1 2 3]\n"

		if got != expected {
			t.Errorf("PrintSlice() = %q, want %q", got, expected)
		}
	})

	t.Run("string_slice", func(t *testing.T) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		PrintSlice([]string{"apple", "banana", "cherry"})

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(r)
		got := buf.String()
		expected := "len=3, cap=3, [apple banana cherry]\n"

		if got != expected {
			t.Errorf("PrintSlice() = %q, want %q", got, expected)
		}
	})

	t.Run("float_slice", func(t *testing.T) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		PrintSlice([]float64{1.1, 2.2, 3.3})

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(r)
		got := buf.String()
		expected := "len=3, cap=3, [1.1 2.2 3.3]\n"

		if got != expected {
			t.Errorf("PrintSlice() = %q, want %q", got, expected)
		}
	})

	t.Run("struct_slice", func(t *testing.T) {
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		people := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}
		PrintSlice(people)

		w.Close()
		os.Stdout = oldStdout

		var buf bytes.Buffer
		buf.ReadFrom(r)
		got := buf.String()
		expected := "len=2, cap=2, [{Alice 25} {Bob 30}]\n"

		if got != expected {
			t.Errorf("PrintSlice() = %q, want %q", got, expected)
		}
	})
}

func TestPrintConvertInfo(t *testing.T) {
	tests := []struct {
		name     string
		from     interface{}
		to       interface{}
		expected string
	}{
		// Basic type conversions
		{
			name:     "integer_to_float",
			from:     func() interface{} { i := 42; return i }(),
			to:       func() interface{} { i := 42; return float64(i) }(),
			expected: "Convert (int, 42) to (float64, 42)\n",
		},
		{
			name: "bool_to_int",
			from: func() interface{} { b := true; return b }(),
			to: func() interface{} {
				b := true
				if b {
					return 1
				} else {
					return 0
				}
			}(),
			expected: "Convert (bool, true) to (int, 1)\n",
		},
		{
			name:     "float64_to_float32",
			from:     func() interface{} { f := 3.14; return f }(),
			to:       func() interface{} { f := 3.14; return float32(f) }(),
			expected: "Convert (float64, 3.14) to (float32, 3.14)\n",
		},

		// Edge cases
		{
			name:     "nil_to_int",
			from:     nil,
			to:       func() interface{} { var i int; return i }(),
			expected: "Convert (<nil>, <nil>) to (int, 0)\n",
		},
		{
			name:     "int_to_string",
			from:     func() interface{} { i := 0; return i }(),
			to:       func() interface{} { i := 0; return fmt.Sprintf("%d", i) }(),
			expected: "Convert (int, 0) to (string, 0)\n",
		},
		{
			name:     "empty_string_to_bytes",
			from:     func() interface{} { s := ""; return s }(),
			to:       func() interface{} { s := ""; return []byte(s) }(),
			expected: "Convert (string, ) to ([]uint8, [])\n",
		},

		// Complex types
		{
			name: "int_slice_to_string_slice",
			from: func() interface{} { nums := []int{1, 2}; return nums }(),
			to: func() interface{} {
				nums := []int{1, 2}
				strs := make([]string, len(nums))
				for i, n := range nums {
					strs[i] = strconv.Itoa(n)
				}
				return strs
			}(),
			expected: "Convert ([]int, [1 2]) to ([]string, [1 2])\n",
		},
		{
			name: "int_map_to_string_map",
			from: func() interface{} { m := map[string]int{"a": 1}; return m }(),
			to: func() interface{} {
				m := map[string]int{"a": 1}
				strMap := make(map[string]string)
				for k, v := range m {
					strMap[k] = strconv.Itoa(v)
				}
				return strMap
			}(),
			expected: "Convert (map[string]int, map[a:1]) to (map[string]string, map[a:1])\n",
		},
		{
			name:     "int_struct_to_string_struct",
			from:     func() interface{} { return struct{ A int }{A: 1} }(),
			to:       func() interface{} { i := 1; return struct{ A string }{A: strconv.Itoa(i)} }(),
			expected: "Convert (struct { A int }, {A:1}) to (struct { A string }, {A:1})\n",
		},

		// Type conversion scenarios
		{
			name:     "int_to_string_conversion",
			from:     func() interface{} { i := 123; return i }(),
			to:       func() interface{} { i := 123; return strconv.Itoa(i) }(),
			expected: "Convert (int, 123) to (string, 123)\n",
		},
		{
			name:     "float_to_int_conversion",
			from:     func() interface{} { f := 3.99; return f }(),
			to:       func() interface{} { f := 3.99; return int(f) }(),
			expected: "Convert (float64, 3.99) to (int, 3)\n",
		},
		{
			name:     "string_to_bytes_conversion",
			from:     func() interface{} { s := "hello"; return s }(),
			to:       func() interface{} { s := "hello"; return []byte(s) }(),
			expected: "Convert (string, hello) to ([]uint8, [104 101 108 108 111])\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Call PrintConvertInfo
			PrintConvertInfo(tt.from, tt.to)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read captured output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			got := buf.String()

			if got != tt.expected {
				t.Errorf("PrintConvertInfo() = %q, want %q", got, tt.expected)
			}
		})
	}
}
