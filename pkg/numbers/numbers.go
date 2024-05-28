package numbers

// SumNumbers sums all the numbers in the provided slice
func SumNumbers(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}
