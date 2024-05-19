package main

import "strconv"

func FizzBuzz(n int) string {

	result := ""

	fizz := map[bool]string{true: "Fizz", false: ""}
	buzz := map[bool]string{true: "Buzz", false: ""}

	fizzPart := fizz[n%3 == 0]
	buzzPart := buzz[n%5 == 0]

	result = fizzPart + buzzPart

	if result == "" {
		result = strconv.Itoa(n)
	}

	return result
}
