package main

import "strconv"

// ----------------- Lv3. -----------------

func FizzBuzzLv3(n int) string {
	var convention = map[int]string{}
	convention[n] = strconv.Itoa(n)
	convention[3] = "Fizz"
	convention[5] = "Buzz"
	return convention[n]
}
