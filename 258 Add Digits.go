func addDigits(num int) int {
	sum := 0
	for num >= 10 {
		sum += num % 10
		num /= 10
		if num < 10 {
			num += sum
			if num < 10 {
				return num
			}
			sum = 0
		}
	}
	return num
}