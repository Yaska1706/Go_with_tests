package integers

import "strconv"

func Add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func Modulus(num string) string {
	numconv, _ := strconv.Atoi(num)
	mod := numconv % 2
	message := "Divisible by 2"

	if mod != 0 {
		message = "Not Divisible by 2"
	}
	return message
}
