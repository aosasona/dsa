package algorithms

import "testing"

var haystack = []string{"Harley", "Bruce", "Clark", "Diana", "Luke", "Lois", "Kara", "Tony", "Pietro"}

func Test_LinearSearchPasses(t *testing.T) {
	needle := "Lois"

	if found := LinearSearch(needle, haystack); !found {
		t.Errorf("Failed to find %s, expected `true`, got `%v`", needle, found)
	}
}
