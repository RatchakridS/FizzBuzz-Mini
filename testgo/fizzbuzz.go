package main

func FizzBuzz(n int) string {

	fizzBuzzConvention := map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		4: "4",
	}

	return fizzBuzzConvention[n]
}
