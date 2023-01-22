package main

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	sumNums := make([]int, 2)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if nums[i]+nums[j] == target {
				sumNums[0] = i
				sumNums[1] = j
				break
			} else {
				continue
			}
		}
	}
	return sumNums
}
