package main

import "testing"

func TestCheckUniqueSymbols(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"aabcd", false},
		{"🙂😀😁", true},
		{"🙂🙂", false},
		{"Привет, мир!", false},
		{"Привет, сад!", true},
	}

	for _, tt := range tests {
		if got := CheckUniqueSymbols(tt.input); got != tt.expected {
			t.Errorf("CheckUniqueSymbols(%q) = %v, expected %v", tt.input, got, tt.expected)
		}
	}
}
