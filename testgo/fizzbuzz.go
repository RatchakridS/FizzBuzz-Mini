package main

import (
	"strconv"
)

// ----------------- Lv1. -----------------

func FizzBuzzLv1(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

// ----------------- Lv2. -----------------

func FizzBuzzLv2(n int) string {
	result := ""
	result += FizzLv2(n)
	result += BuzzLv2(n)
	if result == "" {
		result += strconv.Itoa(n)
	}
	return result
}

func FizzLv2(n int) string {
	mapper := map[bool]string{
		true:  "Fizz",
		false: "",
	}
	return mapper[n%3 == 0]
}

func BuzzLv2(n int) string {
	mapper := map[bool]string{
		true:  "Buzz",
		false: "",
	}
	return mapper[n%5 == 0]
}
