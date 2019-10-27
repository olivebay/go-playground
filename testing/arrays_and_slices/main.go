package main

// Sum adds all the elements of the slice
func Sum(numberToSum []int) int {

	var sums int
	for _, number := range numberToSum {
		sums += number
	}

	return sums
}

// SumAll add all the elemenets of each slice and return the sums of each slice
func SumAll(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails return the last number of the slice
func SumAllTails(numberToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
