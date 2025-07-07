package maximum_number_of_events_that_can_be_attended_test

import (
	"github.com/nelsen129/leetcode-solutions/leetcode/maximum_number_of_events_that_can_be_attended"
	"testing"

	"math/rand/v2"
)

func TestMaxEvents(t *testing.T) {
	tests := []struct {
		Description string
		Events      [][]int
		Want        int
	}{
		{`some overlap`, [][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 4},
		}, 3},
		{`more overlap`, [][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 4},
			[]int{1, 2},
		}, 4},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			got := maximum_number_of_events_that_can_be_attended.MaxEvents(test.Events)
			if got != test.Want {
				t.Errorf(`got %d, want %d`, got, test.Want)
			}
		})
	}
}

func BenchmarkMaxEvents(b *testing.B) {
	sliceSize := 100000
	for b.Loop() {
		data := make([][]int, sliceSize)
		for i := range data {
			num1 := rand.IntN(10000)
			num2 := rand.IntN(10000)
			data[i] = []int{min(num1, num2), max(num1, num2)}
		}

		maximum_number_of_events_that_can_be_attended.MaxEvents(data)
	}
}
