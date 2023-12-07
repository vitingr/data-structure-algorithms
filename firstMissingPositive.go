package main

func firstMissingPositive(nums []int) int {
		// When you need to use constant memorie space you need to SWAP, and HashTables
	for i := 0; i < len(nums); {
		current := nums[i]
		if current > 0 && current <= len(nums) && nums[current - 1] != current {
			nums[i], nums[nums[i] - 1] = nums[nums[i] -1], nums[i]
			continue
		}
		i++
	}

	for i:= 1; i < len(nums) + 1; i++ {
		if nums[i-1] != i {
			return i
		}
	}

	return len(nums) + 1
}

func main() {

}