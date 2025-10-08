package usecase

import (
	"testing"
	"unix_grep_lite/internal/domain"

	"github.com/stretchr/testify/require"
)

func TestWithContext(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		pattern  string
		opts     domain.GrepOptions
		expected string
		wantErr  bool
	}{
		{
			name:     "no context flags",
			input:    "line1\nline2\npattern\nline4\nline5",
			pattern:  "pattern",
			opts:     domain.GrepOptions{},
			expected: "pattern",
		},
		{
			name:    "after context",
			input:   "line1\nline2\npattern\nline4\nline5",
			pattern: "pattern",
			opts: domain.GrepOptions{
				AfterContext: true,
				NumAfter:     2,
			},
			expected: "pattern\nline4\nline5",
		},
		{
			name:    "before context",
			input:   "line1\nline2\npattern\nline4\nline5",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				NumBefore:     2,
			},
			expected: "line1\nline2\npattern",
		},
		{
			name:    "before and after context",
			input:   "line1\nline2\npattern\nline4\nline5",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     1,
				NumAfter:      1,
			},
			expected: "line2\npattern\nline4",
		},
		{
			name:    "around context",
			input:   "line1\nline2\npattern\nline4\nline5",
			pattern: "pattern",
			opts: domain.GrepOptions{
				AroundContext: true,
				NumAround:     2,
			},
			expected: "line1\nline2\npattern\nline4\nline5",
		},
		{
			name:    "multiple matches with separator",
			input:   "line1\npattern\nline3\nline4\nline5\npattern\nline7",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     1,
				NumAfter:      1,
			},
			expected: "line1\npattern\nline3\n--\nline5\npattern\nline7",
		},
		{
			name:    "overlapping contexts",
			input:   "line1\npattern\nline3\npattern\nline5",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     1,
				NumAfter:      1,
			},
			expected: "line1\npattern\nline3\npattern\nline5",
		},
		{
			name:    "with line numbers",
			input:   "line1\npattern\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     1,
				NumAfter:      1,
				LineNumber:    true,
			},
			expected: "1:line1\n2:pattern\n3:line3",
		},
		{
			name:    "context at beginning",
			input:   "pattern\nline2\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     2,
				NumAfter:      1,
			},
			expected: "pattern\nline2",
		},
		{
			name:    "context at end",
			input:   "line1\nline2\npattern",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     1,
				NumAfter:      2,
			},
			expected: "line2\npattern",
		},
		{
			name:    "invert match with context",
			input:   "line1\npattern\nline3\nline4",
			pattern: "pattern",
			opts: domain.GrepOptions{
				InvertMatch:   true,
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     1,
				NumAfter:      1,
			},
			expected: "line1\npattern\nline3\nline4",
		},
		{
			name:     "empty input",
			input:    "",
			pattern:  "pattern",
			opts:     domain.GrepOptions{AroundContext: true, NumAround: 1},
			expected: "",
		},
		{
			name:    "single line match",
			input:   "pattern",
			pattern: "pattern",
			opts: domain.GrepOptions{
				AroundContext: true,
				NumAround:     2,
			},
			expected: "pattern",
		},
		{
			name:    "no matches",
			input:   "line1\nline2\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				AroundContext: true,
				NumAround:     1,
			},
			expected: "",
		},
		{
			name:    "context zero",
			input:   "line1\npattern\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				AfterContext:  true,
				NumBefore:     0,
				NumAfter:      0,
			},
			expected: "pattern",
		},
		{
			name:    "invalid negative before context",
			input:   "line1\npattern\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				BeforeContext: true,
				NumBefore:     -1,
			},
			wantErr: true,
		},
		{
			name:    "invalid negative after context",
			input:   "line1\npattern\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				AfterContext: true,
				NumAfter:     -1,
			},
			wantErr: true,
		},
		{
			name:    "invalid negative around context",
			input:   "line1\npattern\nline3",
			pattern: "pattern",
			opts: domain.GrepOptions{
				AroundContext: true,
				NumAround:     -1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			matcher, err := NewMatcher(tt.pattern, tt.opts)
			require.NoError(t, err, "Failed to create matcher")

			result, err := matcher.withContext(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}
