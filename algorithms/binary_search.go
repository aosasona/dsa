package algorithms

import (
	"math"
)

// A binary checks keeps splitting an ORDERED array down the middle only checking other pair(s) based on whether our pivot (mid point) is larger or smaller than the current pile it is checking
// say we have [0,1,2,3,4,5] and we need to find 5, our mid point is 3, 5 is larger than 3, so we will only be searching [4,5] and we can bother not checking the other half
//
// The time complexity of this algorithm is O(log n) because we are splitting the array into half each time we check amd we are only checking one half
// Again, this algorithm only works on sorted lists, arrays etc...
func BinarySearch(needle int, haystack []int) bool {
	var (
		low  = 0
		high = len(haystack)
	)

	for low < high {
		mid := int(math.Ceil(float64((low + high) / 2)))

		if haystack[mid] == needle {
			return true
		} else if haystack[mid] > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return false
}
