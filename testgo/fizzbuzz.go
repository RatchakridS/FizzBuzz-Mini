package main

import "strconv"

// ----------------- Lv3. -----------------

func FizzBuzzLv3(n int) string {
	convention := map[int]string{}
	convention[n] = strconv.Itoa(n)
	return convention[n]
}
