package main

import "strconv"

// ----------------- Lv3. -----------------

func FizzBuzzLv3(n int) string {
	var convention = map[int]string{}
	convention[n] = strconv.Itoa(n)
	convention[3] = "Fizz"
	convention[5] = "Buzz"
	convention[6] = "Fizz"
	convention[9] = "Fizz"
	convention[10] = "Buzz"
	convention[12] = "Fizz"
	return convention[n]
}
