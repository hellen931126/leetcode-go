var res []string

func letterCombinations(digits string) []string {
	buttons := []string{" ", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	res = []string{}
	if len(digits) > 0 {
		combine(buttons, "", digits, 0)
	}
	return res
}

func combine(buttons []string, str, digits string, index int) {
	if index == len(digits) {
		res = append(res, str)
		return
	}
	keys := buttons[digits[index]-48]
	for _, key := range keys {
		combine(buttons, str+string(key), digits, index+1)
	}

}