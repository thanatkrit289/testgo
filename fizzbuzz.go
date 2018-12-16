package fizzbuzz

import "strconv"

func FizzBuzz(n int) string {
	if n == 3 {
		return "Fizz"
	}

	if n == 5 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
