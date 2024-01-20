package arrays

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}

func SumTail(numbers []int) int {
	tail := numbers[1:]
	sum := 0

	for _, number := range tail {
		sum += number
	}

	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
			continue
		}

		sum = append(sum, SumTail(numbers))
	}

	return sum
}
