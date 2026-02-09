package eval

// indexOfDot returns the index of the first dot in s, or -1 if not found
func indexOfDot(s string) int {
	for i, c := range s {
		if c == '.' {
			return i
		}
	}
	return -1
}
