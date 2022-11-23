package main

func main() {
	tests := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}

	targets := []int{
		9,
		6,
		6,
	}

	for i := 0; i < 3; i++ {
		ret := twoSum(tests[i], targets[i])
		_ = ret
		// DEBUG: fmt.Println(ret)
	}
}

func twoSum(nums []int, target int) []int {
	lenOfNums := len(nums)
	indices := make([]int, 2)

out:
	for i := 0; i < lenOfNums-1; i++ {
		for j := i + 1; j < lenOfNums; j++ {

			sum := nums[i] + nums[j]
			if sum == target {
				indices[0] = i
				indices[1] = j
				break out
			}

		}
	}

	return indices
}
