package usecase

import (
	"strconv"
	"strings"
)

// withoutContext выполняет базовый поиск без контекста (только совпавшие строки)
func (m *Matcher) withoutContext(input string) string {
	if input == "" {
		return ""
	}

	lines := strings.Split(input, "\n")
	matchedLines := []string{}
	for i, line := range lines {
		isMatch := m.lineIsMatch(line)
		// Инверсия результата для флага -v
		if m.opts.InvertMatch {
			isMatch = !isMatch
		}
		if isMatch {
			// Добавление номера строки для флага -n
			if m.opts.LineNumber {
				matchedLines = append(matchedLines, strconv.Itoa(i+1)+":"+line)
			} else {
				matchedLines = append(matchedLines, line)
			}
		}
	}
	return strings.Join(matchedLines, "\n")
}
