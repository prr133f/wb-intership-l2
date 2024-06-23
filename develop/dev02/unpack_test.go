package main

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", ErrInvalidString},
		{"", "", nil},
	}

	for _, test := range tests {
		result, err := UnpackString(test.input)
		if result != test.expected || err != test.err {
			t.Errorf("UnpackString(%q) = %q, %v; want %q, %v", test.input, result, err, test.expected, test.err)
		}
	}
}
