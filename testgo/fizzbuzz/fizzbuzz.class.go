package fizzbuzz

import "strconv"

type FizzBuzz struct {
	Number int
}

func (fz *FizzBuzz) FizzBuzzGame() string {
	result := fizzBuzzCalculate(fz.Number)
	convention := map[int]string{
		0: strconv.Itoa(fz.Number),
		1: "Fizz",
		2: "Buzz",
		3: "FizzBuzz",
	}
	return convention[result]
}

func fizzBuzzCalculate(n int) int {
	if n%15 == 0 {
		return 3
	}
	if n%3 == 0 {
		return 1
	}

	if n%5 == 0 {
		return 2
	}

	return 0
}

func (fz *FizzBuzz) SetNumber(n int) {
	fz.Number = n
}
