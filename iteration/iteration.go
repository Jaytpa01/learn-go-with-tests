package iteration

import "strings"

// Repeat returns `s`, repeated `count` times
func Repeat(s string, count int) string {
	b := strings.Builder{}
	for i := 0; i < count; i++ {
		b.WriteString(s)
	}
	return b.String()
}
