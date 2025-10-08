package usecase

import (
	"testing"
	"unix_grep_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestWithoutContext(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		pattern  string
		opts     domain.GrepOptions
		expected string
	}{
		{
			name:     "empty input",
			input:    "",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "",
		},
		{
			name:     "no matches",
			input:    "hello\nworld\nfoo",
			pattern:  "bar",
			opts:     domain.GrepOptions{},
			expected: "",
		},
		{
			name:     "single match",
			input:    "hello\ntest\nworld",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "multiple matches",
			input:    "test\nhello\ntest\nworld\ntest",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "test\ntest\ntest",
		},
		{
			name:     "regex pattern",
			input:    "test123\nhello\ntest456\nworld",
			pattern:  "test\\d+",
			opts:     domain.GrepOptions{},
			expected: "test123\ntest456",
		},
		{
			name:     "case sensitive",
			input:    "Test\ntest\nTEST",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "ignore case",
			input:    "Test\ntest\nTEST",
			pattern:  "test",
			opts:     domain.GrepOptions{IgnoreCase: true},
			expected: "Test\ntest\nTEST",
		},
		{
			name:     "fixed strings",
			input:    "test.txt\nhello\ntest.log",
			pattern:  "test.",
			opts:     domain.GrepOptions{FixedStrings: true},
			expected: "test.txt\ntest.log",
		},
		{
			name:     "fixed strings vs regex",
			input:    "test1\nhello\ntest.",
			pattern:  "test.",
			opts:     domain.GrepOptions{FixedStrings: false},
			expected: "test1\ntest.",
		},
		{
			name:     "invert match - no original matches",
			input:    "hello\nworld\nfoo",
			pattern:  "bar",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: "hello\nworld\nfoo",
		},
		{
			name:     "invert match - some original matches",
			input:    "hello\ntest\nworld\ntest",
			pattern:  "test",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: "hello\nworld",
		},
		{
			name:     "line numbers",
			input:    "hello\ntest\nworld\ntest",
			pattern:  "test",
			opts:     domain.GrepOptions{LineNumber: true},
			expected: "2:test\n4:test",
		},
		{
			name:     "line numbers with invert",
			input:    "hello\ntest\nworld",
			pattern:  "test",
			opts:     domain.GrepOptions{LineNumber: true, InvertMatch: true},
			expected: "1:hello\n3:world",
		},
		{
			name:     "ignore case with line numbers",
			input:    "Hello\nTEST\nWorld",
			pattern:  "test",
			opts:     domain.GrepOptions{IgnoreCase: true, LineNumber: true},
			expected: "2:TEST",
		},
		{
			name:     "partial matches",
			input:    "testing\nhello\ntester",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "testing\ntester",
		},
		{
			name:     "empty lines",
			input:    "test\n\nhello\n\nworld",
			pattern:  "^$",
			opts:     domain.GrepOptions{},
			expected: "\n",
		},
		{
			name:     "single line input",
			input:    "test",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "single line no match",
			input:    "hello",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "",
		},

		// edge cases
		{
			name:     "newline at end",
			input:    "test\nhello\n",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "multiple newlines",
			input:    "test\n\n\nhello",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "only newlines",
			input:    "\n\n\n",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: "",
		},
		{
			name:     "match empty line",
			input:    "hello\n\nworld",
			pattern:  "^$",
			opts:     domain.GrepOptions{},
			expected: "",
		},
		{
			name:     "match empty line with line numbers",
			input:    "hello\n\nworld",
			pattern:  "^$",
			opts:     domain.GrepOptions{LineNumber: true},
			expected: "2:",
		},
		{
			name:     "invert match with empty lines",
			input:    "hello\n\nworld",
			pattern:  "test",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: "hello\n\nworld",
		},
		{
			name:     "complex regex",
			input:    "user@example.com\ninvalid-email\ntest@domain.org",
			pattern:  "\\w+@\\w+\\.\\w+",
			opts:     domain.GrepOptions{},
			expected: "user@example.com\ntest@domain.org",
		},
		{
			name:     "word boundaries",
			input:    "test\ntesting\ntest123\nmy test",
			pattern:  "\\btest\\b",
			opts:     domain.GrepOptions{},
			expected: "test\nmy test",
		},
		{
			name:     "case insensitive fixed string",
			input:    "Hello World\nhello world\nHELLO WORLD",
			pattern:  "hello",
			opts:     domain.GrepOptions{IgnoreCase: true, FixedStrings: true},
			expected: "Hello World\nhello world\nHELLO WORLD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			matcher, err := NewMatcher(tt.pattern, tt.opts)
			require.NoError(t, err, "Failed to create matcher")

			result := matcher.withoutContext(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestWithoutContextInvalidRegex(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		opts    domain.GrepOptions
	}{
		{
			name:    "invalid regex - unclosed bracket",
			pattern: "[abc",
			opts:    domain.GrepOptions{},
		},
		{
			name:    "invalid regex - unclosed parenthesis",
			pattern: "(test",
			opts:    domain.GrepOptions{},
		},
		{
			name:    "invalid regex - invalid escape",
			pattern: "\\",
			opts:    domain.GrepOptions{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewMatcher(tt.pattern, tt.opts)
			require.Error(t, err, "Expected error for invalid regex pattern")
		})
	}
}
