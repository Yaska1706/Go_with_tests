package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numberstoSum ...[]int) []int {
	var sums []int

	for _, number := range numberstoSum {
		sums = append(sums, Sum(number))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) int {
	var sums []int
	sum := 0

	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
			sum = Sum(sums)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
			sum = Sum(sums)
		}

	}
	return sum
}
