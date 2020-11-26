// Problem - https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for k, l := range nums[i+1:] {
			if j+l == target {
				return []int{i, k + i + 1}
			}
		}
	}
	return nil
}