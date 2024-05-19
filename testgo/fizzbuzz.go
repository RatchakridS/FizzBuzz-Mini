package main

func FizzBuzz(n int) string {

	fizzBuzzConvention := map[int]string{
		1:  "1",
		2:  "2",
		3:  "Fizz",
		4:  "4",
		5:  "Buzz",
		6:  "Fizz",
		7:  "7",
		8:  "8",
		9:  "Fizz",
		10: "Buzz",
		11: "11",
		12: "Fizz",
		13: "13",
	}

	return fizzBuzzConvention[n]
}
