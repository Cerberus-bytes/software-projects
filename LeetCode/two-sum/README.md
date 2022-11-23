# Two Sum (Easy)

## Problem Description
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

Example 2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

Example 3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
```

Constraints:
- 2 <= nums.length <= 104
- -109 <= nums[i] <= 109
- -109 <= target <= 109
- Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

## Statistics
### Golang
```
for i in {1..10}; do time ./two_sum; sleep 2; done
./two_sum  0.01s user 0.01s system 69% cpu 0.017 total
./two_sum  0.01s user 0.00s system 67% cpu 0.016 total
./two_sum  0.00s user 0.01s system 61% cpu 0.017 total
./two_sum  0.01s user 0.01s system 67% cpu 0.015 total
./two_sum  0.01s user 0.01s system 73% cpu 0.016 total
./two_sum  0.01s user 0.01s system 77% cpu 0.027 total
./two_sum  0.01s user 0.01s system 75% cpu 0.019 total
./two_sum  0.01s user 0.01s system 77% cpu 0.014 total
./two_sum  0.01s user 0.01s system 75% cpu 0.018 total
./two_sum  0.00s user 0.01s system 64% cpu 0.016 total
```