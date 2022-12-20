package main

import (
    "fmt"
)

func main() {
    nums := []int{0, 3, 1}
    missingNum := missingNumber(nums)
    fmt.Println(missingNum)
}

func missingNumber(nums []int) int {
    sortedSlice := selectionSort(nums)
    
    for i, num := range sortedSlice {
        if i != num {
            return i
        }
    }

    return sortedSlice[len(sortedSlice) - 1] + 1
}

func selectionSort(nums []int) []int {
    lenOfSlice := len(nums)

    for i := 0; i < lenOfSlice - 1; i++ {
        for j := i + 1; j < lenOfSlice; j++ {
            if nums[j] < nums[i] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }

    return nums
}