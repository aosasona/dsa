package algorithms

import "testing"

type BinarySearchTest struct {
	needle     int
	shouldFind bool
}

func Test_BinarySearch(t *testing.T) {
	haystack := []int{20, 24, 52, 89, 101, 134, 174, 258, 267, 280}

	tests := []BinarySearchTest{
		{
			needle:     20,
			shouldFind: true,
		},
		{
			needle:     24,
			shouldFind: true,
		},
		{
			needle:     258,
			shouldFind: true,
		},
		{
			needle:     80,
			shouldFind: false,
		},
	}

	for _, test := range tests {
		if found := BinarySearch(test.needle, haystack); found != test.shouldFind {
			t.Errorf("Failed to find %d, expected `true`, got `%v`", test.needle, found)
		}
	}
}
