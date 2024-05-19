package main

import (
	"strconv"
)

var fizz = "Fizz"
var buzz = "Buzz"

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

func FizzBuzzImprove(n int) string {
	value := strconv.Itoa(n)
	n = RollingNumberInScope(n)
	fizzBuzzConvention := map[int]string{
		1:  value,
		2:  value,
		3:  fizz,
		4:  value,
		5:  buzz,
		6:  fizz,
		7:  value,
		8:  value,
		9:  fizz,
		10: buzz,
		11: value,
		12: fizz,
		13: value,
		14: value,
		15: fizz + buzz,
	}

	return fizzBuzzConvention[n]
}

func RollingNumberInScope(n int) int {
	if n <= 15 {
		return n
	}
	value := n % 15
	if value == 0 {
		return 15
	}
	return value
}

// func Level3(n int) string {

// }
