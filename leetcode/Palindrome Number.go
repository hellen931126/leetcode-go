import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i, n := 0, length/2; i < n; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}