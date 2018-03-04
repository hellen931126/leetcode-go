func romanToInt(s string) int {
	romanToIntTable := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	exNum, res := 0, 0
	for i := len(s); i > 0; i-- {
		roman := string([]rune(s)[i-1])
		num := romanToIntTable[roman]
		if exNum > num {
			res, exNum = res-num, num
			continue
		}
		res, exNum = res+num, num

	}
	return res
}