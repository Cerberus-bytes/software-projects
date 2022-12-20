package main

func main() {
	test1 := [-1,0,3,5,9,12]
	target1 := 9
	index := search(test1, target1)
}

func search(nums []int, target int) int {
    var m int
    h := len(nums) - 1
    l := 0

    for l <= h {
        m = (h + l) / 2
        fmt.Println(m)
        if nums[m] == target {
            return m
        } else if nums[m] < target {
            l = m + 1
        } else if nums[m] > target {
            h = m - 1
        }
    }

    return -1
}