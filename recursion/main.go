package main

// Count number of steps to reduce a number to zero
// in decrement of 1
func ReduceToZero(counter *int, number int) int {
	if number < 1 {
		return number
	}
	number = number - 1
	*counter = *counter + 1
	return ReduceToZero(counter, number)
}
