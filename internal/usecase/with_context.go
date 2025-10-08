package usecase

import (
	"strconv"
	"strings"
	"unix_grep_lite/internal/domain"
)

// Line структура строки с её содержимым и номером
type Line struct {
	val string
	num int // номер строки (начиная с 1)
}

// withContext обрабатывает поиск с контекстом (строки до/после совпадений)
func (m *Matcher) withContext(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	if m.opts.NumAfter < 0 || m.opts.NumBefore < 0 || m.opts.NumAround < 0 {
		return "", domain.ErrInvalidContextLength
	}

	// Преобразование -C в -A и -B, т.к. -AB 1 ~ -C 1
	if m.opts.AroundContext {
		m.opts.BeforeContext, m.opts.AfterContext = true, true
		m.opts.NumBefore, m.opts.NumAfter = m.opts.NumAround, m.opts.NumAround
	}

	beforeN, afterN := m.opts.NumBefore, m.opts.NumAfter
	seens := make(map[int]bool)     // предотвращает дублирование строк
	matchedLines := make([]Line, 0) // результирующие строки с номерами
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		isMatch := m.lineIsMatch(line)
		// Инверсия результата для флага -v
		if m.opts.InvertMatch {
			isMatch = !isMatch
		}
		if isMatch {
			// Добавление контекстных строк до совпадения
			if m.opts.BeforeContext {
				for j := max(i-beforeN, 0); j < i; j++ {
					if seens[j] {
						continue
					}
					seens[j] = true
					matchedLines = append(matchedLines, Line{val: lines[j], num: j + 1})
				}
			}
			// Добавление самой совпавшей строки
			if !seens[i] {
				matchedLines = append(matchedLines, Line{val: lines[i], num: i + 1})
				seens[i] = true
			}
			// Добавление контекстных строк после совпадения
			if m.opts.AfterContext {
				for j := i + 1; j <= min(i+afterN, len(lines)-1); j++ {
					if seens[j] {
						continue
					}
					seens[j] = true
					matchedLines = append(matchedLines, Line{val: lines[j], num: j + 1})
				}
			}
		}
	}

	result := joinLinesContextsWithSep(matchedLines, "--", m.opts.LineNumber)
	return strings.Join(result, "\n"), nil
}

// joinLinesContextsWithSep объединяет строки с разделителем между группами
func joinLinesContextsWithSep(matchedLines []Line, sep string, isLineNumber bool) []string {
	result := make([]string, 0, len(matchedLines))
	for i := 0; i < len(matchedLines); i++ {
		// Добавление номера строки для флага -n
		if isLineNumber {
			result = append(result, strconv.Itoa(matchedLines[i].num)+":"+matchedLines[i].val)
		} else {
			result = append(result, matchedLines[i].val)
		}
		// Вставка разделителя между несмежными группами строк
		if i+1 < len(matchedLines) && matchedLines[i+1].num-matchedLines[i].num > 1 {
			result = append(result, sep)
		}
	}
	return result
}
