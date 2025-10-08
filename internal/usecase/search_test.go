package usecase

import (
	"testing"
	"unix_grep_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestSearchMatch(t *testing.T) {
	tests := []struct {
		name           string
		pattern        string
		input          string
		opts           domain.GrepOptions
		expected       string
		wantCompileErr bool
		wantSearchErr  bool
	}{
		{
			name:     "basic match",
			pattern:  "test",
			input:    "hello\ntest\nworld",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "no matches",
			pattern:  "missing",
			input:    "hello\ntest\nworld",
			opts:     domain.GrepOptions{},
			expected: "",
		},
		{
			name:     "count mode",
			pattern:  "test",
			input:    "test\nhello\ntest\nworld",
			opts:     domain.GrepOptions{Count: true},
			expected: "2",
		},
		{
			name:     "with line numbers",
			pattern:  "test",
			input:    "hello\ntest\nworld",
			opts:     domain.GrepOptions{LineNumber: true},
			expected: "2:test",
		},
		{
			name:    "after context",
			pattern: "test",
			input:   "hello\ntest\nworld\nfoo",
			opts: domain.GrepOptions{
				AfterContext: true,
				NumAfter:     1,
			},
			expected: "test\nworld",
		},
		{
			name:    "before context",
			pattern: "test",
			input:   "hello\ntest\nworld",
			opts: domain.GrepOptions{
				BeforeContext: true,
				NumBefore:     1,
			},
			expected: "hello\ntest",
		},
		{
			name:    "around context",
			pattern: "test",
			input:   "hello\ntest\nworld",
			opts: domain.GrepOptions{
				AroundContext: true,
				NumAround:     1,
			},
			expected: "hello\ntest\nworld",
		},
		{
			name:     "ignore case",
			pattern:  "TEST",
			input:    "hello\ntest\nworld",
			opts:     domain.GrepOptions{IgnoreCase: true},
			expected: "test",
		},
		{
			name:     "fixed strings",
			pattern:  "test.",
			input:    "hello\ntest.\nworld",
			opts:     domain.GrepOptions{FixedStrings: true},
			expected: "test.",
		},
		{
			name:     "invert match",
			pattern:  "test",
			input:    "hello\ntest\nworld",
			opts:     domain.GrepOptions{InvertMatch: true},
			expected: "hello\nworld",
		},
		{
			name:     "regex pattern",
			pattern:  "t.st",
			input:    "hello\ntest\nworld",
			opts:     domain.GrepOptions{},
			expected: "test",
		},
		{
			name:     "empty input",
			pattern:  "test",
			input:    "",
			opts:     domain.GrepOptions{},
			expected: "",
		},
		{
			name:           "invalid regex",
			pattern:        "[",
			input:          "test",
			opts:           domain.GrepOptions{},
			wantCompileErr: true,
		},
		{
			name:    "context error",
			pattern: "test",
			input:   "test",
			opts: domain.GrepOptions{
				AfterContext: true,
				NumAfter:     -1,
			},
			wantSearchErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			matcher, err := NewMatcher(tt.pattern, tt.opts)
			if tt.wantCompileErr {
				require.Error(t, err)
				return
			}

			result, err := matcher.SearchMatch(tt.pattern, tt.input, tt.opts)
			if tt.wantSearchErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}
