package maximum_number_of_events_that_can_be_attended

import (
	"container/heap"
	"sort"
)

// Probem 1353. https://leetcode.com/problems/maximum-number-events-that-can-be-attended
// You are given an array of events where events[i] = [startDay_i, endDay_i]. Every event i starts at startDay_i and ends at endDay_i.
// You can attend an event i at any day d where startTime_i <= d <= endTime_i. You can only attend one event at any time d.
// Return the maximum number of events you can attend.
func MaxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	var eventIdx, attended, day int
	validEvents := IntHeap{}
	heap.Init(&validEvents)

	for eventIdx < len(events) || len(validEvents) > 0 {
		for eventIdx < len(events) && events[eventIdx][0] <= day {
			heap.Push(&validEvents, events[eventIdx][1])
			eventIdx++
		}

		// get rid of outdated events
		for len(validEvents) > 0 && validEvents[0] < day {
			heap.Pop(&validEvents)
		}

		if len(validEvents) > 0 && validEvents[0] >= day {
			attended++
			heap.Pop(&validEvents)
		}

		if len(validEvents) > 0 {
			day = day + 1
		} else if eventIdx < len(events) {
			day = events[eventIdx][0]
		}
	}

	return attended
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}
