package main

import (
	"fmt"
	"io"
	"os"
	"unix_grep_lite/internal/domain"
	"unix_grep_lite/internal/usecase"

	"github.com/spf13/pflag"
)

func main() {
	// flags init
	numAfter := pflag.IntP("after-context", "A", 0, "Print num lines of trailing context after matching lines.")
	numBefore := pflag.IntP("before-context", "B", 0, "Print num lines of leading context before matching lines.")
	numAround := pflag.IntP("context", "C", 0, "Print num lines of leading and trailing output context.")
	count := pflag.BoolP("count", "c", false, "Suppress normal output; instead print a count of matching lines for each input file.")
	ignoreCase := pflag.BoolP("ignore-case", "i", false, "Ignore case distinctions in patterns and input data, so that characters that differ only in case match each other.")
	invertMatch := pflag.BoolP("invert-match", "v", false, "Invert the sense of matching, to select non-matching lines.")
	fixedStrings := pflag.BoolP("fixed-strings", "F", false, "Interpret patterns as fixed strings, not regular expressions.")
	lineNumber := pflag.BoolP("line-number", "n", false, "Prefix each line of output with the 1-based line number within its input file.")

	pflag.Parse()

	opts := domain.GrepOptions{
		NumAfter:     *numAfter,
		NumBefore:    *numBefore,
		NumAround:    *numAround,
		Count:        *count,
		IgnoreCase:   *ignoreCase,
		InvertMatch:  *invertMatch,
		FixedStrings: *fixedStrings,
		LineNumber:   *lineNumber,
	}
	pflag.Visit(func(f *pflag.Flag) {
		if f.Name == "after-context" {
			opts.AfterContext = true
		}
		if f.Name == "before-context" {
			opts.BeforeContext = true
		}
		if f.Name == "context" {
			opts.AroundContext = true
		}
	})

	args := pflag.Args()

	var (
		pattern, input string
		b              []byte
		err            error
	)
	switch {
	case len(args) == 1:
		// Читаем из stdin если файлы не указаны
		b, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	case len(args) == 2:
		// Читаем из файла
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		defer file.Close() //nolint:errcheck

		b, err = io.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Error:", domain.ErrWrongArgs)
		os.Exit(1)
	}

	pattern = args[0]
	input = string(b)

	matcher, err := usecase.NewMatcher(pattern, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create matcher: %w", err)
		os.Exit(1)
	}
	result, err := matcher.SearchMatch(pattern, input, opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
