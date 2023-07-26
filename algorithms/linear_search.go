package algorithms

func LinearSearch(needle string, haystack []string) bool {
	var (
		i = 0

		found = false
	)

	for i < len(haystack) {
		if haystack[i] == needle {
			found = true
			break
		}
		i++
	}

	return found
}
