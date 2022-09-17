package repeat

import "strings"

// Repeat repeats s n times
func Repeat(s string, n int) string {
	var repeat string
	for i := 0; i < n; i++ {
		repeat += s
	}

	return repeat
}

// RepeatSl repeats s n times
func RepeatSl(s string, n int) string {
	repeat := make([]string, 0, n)
	for i := 0; i < n; i++ {
		repeat = append(repeat, s)
	}

	return strings.Join(repeat, ",")
}
