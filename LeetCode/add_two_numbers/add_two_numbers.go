package main

import(

)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    // Test1: l1 = [2,4,3], l2 = [5,6,4], returns [7,0,8]
    // Test2: l1 = [0], l2 = [0], return [0]
    // Test3: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9], returns [8,9,9,9,0,0,0,1]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var sum int    
    var onesPlaceDigit int
    var isCarry bool
    var headPtr *ListNode

    for l1 != nil || l2 != nil || isCarry {
        sum = addVals(l1, l2, isCarry)
        if sum == -1 {
            panic("Error: Could not add")
        }

        isCarry = carryCheck(sum)
        onesPlaceDigit = getOnes(sum)
        headPtr = linkNode(headPtr, onesPlaceDigit)
        
        l1 = nextNode(l1)
        l2 = nextNode(l2)
    }

    return headPtr
}

func nextNode(ln *ListNode) *ListNode {
    if ln != nil {
        return ln.Next
    }
    return nil
}

func addVals(l1 *ListNode, l2 *ListNode, isCarry bool) int {
    if l1 != nil && l2 != nil && isCarry {
        return l1.Val + l2.Val + 1
    } else if l1 != nil && l2 == nil && isCarry {
        return l1.Val + 1
    } else if l1 == nil && l2 != nil && isCarry {
        return l2.Val + 1
    } else if l1 == nil && l2 == nil && isCarry {
        return 1
    } else if l1 != nil && l2 != nil && !isCarry {
        return l1.Val + l2.Val
    } else if l1 != nil && l2 == nil && !isCarry {
        return l1.Val
    } else if l1 == nil && l2 != nil && !isCarry {
        return l2.Val
    }

    return -1
}

func linkNode(headPtr *ListNode, digit int) *ListNode {
    newNode := &ListNode{
        Val: digit,
        Next: nil,
    }

    if headPtr == nil {
        headPtr = newNode
        return headPtr
    }

    tmpPtr := headPtr
    for tmpPtr.Next != nil {
        tmpPtr = tmpPtr.Next
    }
    tmpPtr.Next = newNode

    return headPtr
}


func getOnes(sum int) int {
    strSum := strconv.Itoa(sum)
    onesPlaceDigitStr := string(strSum[len(strSum) - 1])
    onesPlaceDigitInt, err := strconv.Atoi(onesPlaceDigitStr)
    if err != nil {
        panic(err)
    }
    
    return onesPlaceDigitInt
}

func carryCheck(sum int) bool { 
    strSum := strconv.Itoa(sum)
    
    if len(strSum) > 1 {
        return true
    }

    return false
}