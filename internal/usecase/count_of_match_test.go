package usecase

import (
	"testing"
	"unix_grep_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestCountOfMatching(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		pattern  string
		opts     domain.GrepOptions
		expected int
	}{
		{
			name:     "empty input",
			input:    "",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 0,
		},
		{
			name:     "no matches",
			input:    "hello\nworld\nfoo",
			pattern:  "bar",
			opts:     domain.GrepOptions{},
			expected: 0,
		},
		{
			name:     "single match",
			input:    "hello\ntest\nworld",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 1,
		},
		{
			name:     "multiple matches",
			input:    "test\nhello\ntest\nworld\ntest",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 3,
		},
		{
			name:     "regex pattern",
			input:    "test123\nhello\ntest456\nworld",
			pattern:  "test\\d+",
			opts:     domain.GrepOptions{},
			expected: 2,
		},
		{
			name:     "case sensitive",
			input:    "Test\ntest\nTEST",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 1,
		},
		{
			name:     "ignore case",
			input:    "Test\ntest\nTEST",
			pattern:  "test",
			opts:     domain.GrepOptions{IgnoreCase: true},
			expected: 3,
		},
		{
			name:     "fixed strings",
			input:    "test.txt\nhello\ntest.log",
			pattern:  "test.",
			opts:     domain.GrepOptions{FixedStrings: true},
			expected: 2,
		},
		{
			name:     "fixed strings vs regex",
			input:    "test.txt\nhello\ntest.log",
			pattern:  "test.",
			opts:     domain.GrepOptions{FixedStrings: false},
			expected: 2,
		},
		{
			name:     "invert match - no matches",
			input:    "hello\nworld\nfoo",
			pattern:  "bar",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: 3,
		},
		{
			name:     "invert match - some matches",
			input:    "hello\ntest\nworld\ntest",
			pattern:  "test",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: 2,
		},
		{
			name:     "invert match with ignore case",
			input:    "Hello\nTEST\nWorld",
			pattern:  "test",
			opts:     domain.GrepOptions{InvertMatch: true, IgnoreCase: true},
			expected: 2,
		},
		{
			name:     "single line input",
			input:    "test",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 1,
		},
		{
			name:     "empty lines",
			input:    "test\n\nhello\n\nworld",
			pattern:  "^$",
			opts:     domain.GrepOptions{},
			expected: 2,
		},
		{
			name:     "partial matches",
			input:    "testing\nhello\ntester",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 2,
		},

		// edge cases
		{
			name:     "newline at end",
			input:    "test\nhello\n",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 1,
		},
		{
			name:     "multiple newlines",
			input:    "test\n\n\nhello",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 1,
		},
		{
			name:     "only newlines",
			input:    "\n\n\n",
			pattern:  "test",
			opts:     domain.GrepOptions{},
			expected: 0,
		},
		{
			name:     "match empty line with invert",
			input:    "hello\n\nworld",
			pattern:  "test",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			matcher, err := NewMatcher(tt.pattern, tt.opts)
			require.NoError(t, err, "Failed to create matcher")

			result := matcher.countOfMatching(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
