package longest_increasing_subsequence_test

import (
	"github.com/nelsen129/leetcode-solutions/leetcode/longest_increasing_subsequence"
	"testing"

	"math/rand/v2"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		Description string
		Nums        []int
		Want        int
	}{
		{`mixed unique`, []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{`some duplicates`, []int{0, 1, 0, 3, 2, 3}, 4},
		{`all duplicates`, []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{`mixes bag`, []int{4, 10, 4, 3, 8, 9}, 3},
		{`sorted`, []int{1, 2, 3, 4, 5, 6}, 6},
		{`reverse sorted`, []int{6, 5, 4, 3, 2, 1}, 1},
		{`empty`, []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			got := longest_increasing_subsequence.LengthOfLIS(test.Nums)
			if got != test.Want {
				t.Errorf(`got %d, want %d`, got, test.Want)
			}
		})
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	sliceSize := 100000
	for b.Loop() {
		data := make([]int, sliceSize)
		for i := range data {
			// data is in range -10e4 through 10e4
			data[i] = rand.IntN(20000) - 10000
		}

		longest_increasing_subsequence.LengthOfLIS(data)
	}
}
