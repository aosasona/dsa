package algorithms

import "testing"

func Test_BubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{21, 10, 16, 4, 7, 9, 1}, []int{1, 4, 7, 9, 10, 16, 21}},
		{[]int{33, 19, 26, 18, 5, 9}, []int{5, 9, 18, 19, 26, 33}},
	}

	for _, test := range tests {
		if got := BubbleSort(test.input); !Equal(got, test.expected) {
			t.Errorf("BubbleSort(%v) = %v", test.input, got)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, val := range a {
		if val != b[i] {
			return false
		}
	}

	return true
}

func Benchmark_BubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{2, 3, 1})
	}
}
