package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"главрыба", "абырвалг"},
		{"hello World", "dlroW olleh"},
		{"latin кириллица", "ациллирик nital"},
		{"", ""},
		{"a", "a"},
		{"абвгдеЕёЁ", "ЁёЕедгвба"},
		{"🦋Golang🐍", "🐍gnaloG🦋"},
	}

	for _, tt := range tests {
		result := reverseString(tt.input)
		if result != tt.expected {
			t.Errorf("reverseString(%q) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}
