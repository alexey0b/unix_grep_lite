package usecase

import (
	"strings"
)

// countOfMatching подсчитывает количество совпавших строк (флаг -c)
func (m *Matcher) countOfMatching(input string) int {
	if input == "" {
		return 0
	}

	cnt := 0
	lines := strings.SplitSeq(input, "\n") // итератор по строкам
	for line := range lines {
		isMatch := m.lineIsMatch(line)
		// Инверсия для флага -v
		if m.opts.InvertMatch {
			isMatch = !isMatch
		}
		if isMatch {
			cnt++
		}
	}
	return cnt
}
