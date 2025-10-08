package usecase

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unix_grep_lite/internal/domain"
)

// Matcher структура для поиска с предкомпилированным паттерном
type Matcher struct {
	pattern       string         // для фиксированных строк
	compiledRegex *regexp.Regexp // для регулярных выражений
	opts          domain.GrepOptions
}

// NewMatcher создает matcher с с
func NewMatcher(pattern string, opts domain.GrepOptions) (*Matcher, error) {
	m := &Matcher{opts: opts}

	// Обработка фиксированных строк и регулярных выражений
	if opts.FixedStrings {
		if opts.IgnoreCase {
			m.pattern = strings.ToLower(pattern)
		} else {
			m.pattern = pattern
		}
	} else {
		regexPattern := pattern
		if opts.IgnoreCase {
			regexPattern = "(?i)" + pattern
		}

		var err error
		m.compiledRegex, err = regexp.Compile(regexPattern)
		if err != nil {
			return nil, fmt.Errorf("invalid regexp pattern '%s': %w", pattern, err)
		}
	}

	return m, nil
}

// SearchMatch выполняет поиск паттерна в тексте с заданными опциями
func (m *Matcher) SearchMatch(pattern, input string, opts domain.GrepOptions) (string, error) {
	// Выбор режима обработки на основе опций
	var (
		result string
		err    error
	)
	switch {
	case opts.Count:
		result = strconv.Itoa(m.countOfMatching(input))
	case opts.AfterContext || opts.BeforeContext || opts.AroundContext:
		result, err = m.withContext(input)
		if err != nil {
			return "", fmt.Errorf("context processing failed: %w", err)
		}
	default:
		result = m.withoutContext(input)
	}
	return result, nil
}

// lineIsMatch проверяет соответствие строки паттерну
func (m *Matcher) lineIsMatch(line string) bool {
	if m.opts.IgnoreCase {
		line = strings.ToLower(line)
	}
	if m.opts.FixedStrings {
		return strings.Contains(line, m.pattern)
	}
	return m.compiledRegex.MatchString(line)
}
