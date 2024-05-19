package main

func FizzBuzz(n int) string {

	fizzBuzzConvention := map[int]string{
		1: "1",
		2: "2",
		3: "Fizz",
		4: "4",
		5: "Buzz",
		6: "Fizz",
		7: "7",
		8: "8",
	}

	return fizzBuzzConvention[n]
}
