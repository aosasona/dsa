package algorithms

import "testing"

func Test_LinearSearch(t *testing.T) {
	haystack := []string{"Harley", "Bruce", "Clark", "Diana", "Luke", "Lois", "Kara", "Tony", "Pietro"}
	needle := "Lois"

	if found := LinearSearch(needle, haystack); !found {
		t.Errorf("Failed to find %s, expected `true`, got `%v`", needle, found)
	}
}
