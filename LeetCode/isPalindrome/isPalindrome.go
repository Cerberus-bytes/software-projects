func main() {
    palindrome := isPalindrome(5)
}

func isPalindrome(x int) bool {
    originalStr := strconv.Itoa(x)
    reverseStr := reverse(originalStr)
    
    var isPalindrome bool

    if originalStr == reverseStr {
        isPalindrome = true
    }

    return isPalindrome
}

func reverse(originalStr string) string {
    strLen := len(originalStr) - 1
    reverseStr := ""

    for i := strLen ; i > -1 ; i-- {
        reverseStr = reverseStr + string(originalStr[i])
    }

    return reverseStr
}