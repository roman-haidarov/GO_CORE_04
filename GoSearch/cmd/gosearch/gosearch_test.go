package main

import "testing"

func Test_containsWord(t *testing.T) {
	tests := []struct {
		text     string
		word     string
		expected bool
	}{
		{
			text:     "docs",
			word:     "test",
			expected: true,
		},
		{
			text:     "Hello, world!",
			word:     "Go",
			expected: false,
		},
	}

	for _, tt := range tests {
		result := containsWord(tt.text, tt.word)
		if result != tt.expected {
			t.Errorf("containsWord(%q, %q) = %v, expected %v", tt.text, tt.word, result, tt.expected)
		}
	}
}
