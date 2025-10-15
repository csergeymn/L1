package main

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"English words", "snow dog sun", "sun dog snow"},
		{"Russian words", "снег собака солнце", "солнце собака снег"},
		{"Single word", "привет", "привет"},
		{"Two words", "hello world", "world hello"},
		{"With emoji", "cat 🐱 dog 🐶", "🐶 dog 🐱 cat"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.input)

			if got != tt.expected {
				t.Errorf("reverseWords(%q) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
