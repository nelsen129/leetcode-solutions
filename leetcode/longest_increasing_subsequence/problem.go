package longest_increasing_subsequence

// Probem 300. https://leetcode.com/problems/longest-increasing-subsequence
// Given an integer array nums, return the length of the
// longest strictly increasing subsequence
func LengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// longest contains the minimal increasing subsequence of every length
	longest := [][]int{[]int{nums[0]}}

	for i := 1; i < len(nums); i++ {
		for j, l := range longest {
			if nums[i] < l[len(l)-1] { // is nums[i] a better pick for the subsequence?
				if (len(l) > 1 && nums[i] > l[len(l)-2]) || len(l) == 1 {
					l[len(l)-1] = nums[i]
				}
			} else if nums[i] > l[len(l)-1] { // is nums[i] a good pick for the next subsequence?
				if j < len(longest)-1 && nums[i] < longest[j+1][len(longest[j+1])-1] {
					longest[j+1] = append(l, nums[i])
				} else if j == len(longest)-1 {
					longest = append(longest, append(l, nums[i]))
					//break
				}
			}
		}
	}

	return len(longest[len(longest)-1])
}
