package annotations

import "strings"

func CleanContent(v string) string { return strings.TrimSpace(strings.ReplaceAll(v, "\n", " ")) }
func IsSubstantive(v string) bool  { return len(CleanContent(v)) >= 3 }
func PageCount(total, size int) int {
	if size <= 0 {
		return 0
	}
	n := total / size
	if total%size != 0 {
		n++
	}
	return n
}
