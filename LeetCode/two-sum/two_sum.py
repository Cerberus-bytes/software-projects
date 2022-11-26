#!/usr/bin/env python3

def twoSum(nums, target):
    lenOfList = len(nums)
    for i in range(0, lenOfList - 1):
        for j in range(i + 1, lenOfList):
            
            sum = nums[i] + nums[j]
            if sum == target:
                return [i, j]

nums = [2,7,11,15]
target = 9

soln = twoSum(nums, target)