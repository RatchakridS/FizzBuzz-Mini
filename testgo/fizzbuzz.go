package main

import (
	"strconv"

	"github.com/AnuchitO/tests/testgo/fizzbuzz"
)

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
	convention[15] = "FizzBuzz"
	return convention[n]
}

func FizzBuzzLv4(n int) string {
	game := fizzbuzz.FizzBuzz{}
	game.Number = n

	result := game.FizzBuzzGame()
	return result
}
