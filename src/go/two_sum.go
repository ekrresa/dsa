package main

func TwoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		index, found := store[diff]

		if !found {
			store[num] = i
		} else {
			return []int{index, i}
		}
	}

	return []int{}
}
