package analyzer

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func ReplaceWith(s, old string, news []string) string {
	if len(news) == 0 {
		return s
	}
	r := []rune(strconv.Quote(s))
	s = string(r[1 : len(r)-1]) // Remove enclosing quotes.
	// Compute number of replacements.
	n := strings.Count(s, old)
	if n == 0 {
		return s
	}
	n = min(n, len(news))
	var b strings.Builder
	start := 0
	if len(old) > 0 {
		for i := range n {
			j := start + strings.Index(s[start:], old)
			if j-start > 1 {
				b.WriteString(` + "`)
			}
			b.WriteString(s[start:j])
			if j-start > 1 {
				b.WriteString(`" + `)
			}
			b.WriteString(news[i])
			start = j + len(old)
		}
	} else { // len(old) == 0
		b.WriteString(news[0])
		for i := range n - 1 {
			b.WriteString(" + ")
			_, wid := utf8.DecodeRuneInString(s[start:])
			j := start + wid
			b.WriteString(s[start:j])
			b.WriteString(news[i+1])
			start = j
		}
	}
	b.WriteString(s[start:])
	return b.String()
}
