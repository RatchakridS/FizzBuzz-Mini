package main

import (
	"fmt"
	"testing"
)

func TestLv1(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
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
		14: "14",
		15: "FizzBuzz",
	}
	for _, v := range inputList {
		actual := FizzBuzzLv1(v)
		expect := fizzBuzzConvention[v]
		t.Run(fmt.Sprintf("%s%s%s%d", "lv1:ShouldReturn", expect, "WhenInput", v), func(t *testing.T) {
			if expect != actual {
				t.Errorf("got %q but want %q", actual, expect)
			}
		})
	}
}

func TestLv2(t *testing.T) {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
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
		14: "14",
		15: "FizzBuzz",
	}
	for _, v := range inputList {
		actual := FizzBuzzLv2(v)
		expect := fizzBuzzConvention[v]
		t.Run(fmt.Sprintf("%s%s%s%d", "lv1:ShouldReturn", expect, "WhenInput", v), func(t *testing.T) {
			if expect != actual {
				t.Errorf("got %q but want %q", actual, expect)
			}
		})
	}
}

// func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
// 	input := 1

// 	got := FizzBuzzImprove(input)

// 	want := "1"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
// 	input := 2

// 	got := FizzBuzzImprove(input)

// 	want := "2"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzWhenInput3(t *testing.T) {
// 	input := 3

// 	got := FizzBuzzImprove(input)

// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn4WhenInput4(t *testing.T) {
// 	input := 4

// 	got := FizzBuzzImprove(input)

// 	want := "4"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnBuzzWhenInput5(t *testing.T) {
// 	input := 5

// 	got := FizzBuzzImprove(input)

// 	want := "Buzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzWhenInput6(t *testing.T) {
// 	input := 6

// 	got := FizzBuzzImprove(input)

// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn7WhenInput7(t *testing.T) {
// 	input := 7

// 	got := FizzBuzzImprove(input)

// 	want := "7"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn8WhenInput8(t *testing.T) {
// 	input := 8

// 	got := FizzBuzzImprove(input)

// 	want := "8"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn9WhenInput9(t *testing.T) {
// 	input := 9

// 	got := FizzBuzzImprove(input)

// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnBuzzWhenInput10(t *testing.T) {
// 	input := 10

// 	got := FizzBuzzImprove(input)

// 	want := "Buzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn11WhenInput11(t *testing.T) {
// 	input := 11

// 	got := FizzBuzzImprove(input)

// 	want := "11"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzWhenInput12(t *testing.T) {
// 	input := 12

// 	got := FizzBuzzImprove(input)

// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn13WhenInput13(t *testing.T) {
// 	input := 13

// 	got := FizzBuzzImprove(input)

// 	want := "13"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn14WhenInput14(t *testing.T) {
// 	input := 14

// 	got := FizzBuzzImprove(input)

// 	want := "14"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzBuzzWhenInput15(t *testing.T) {
// 	input := 15

// 	got := FizzBuzzImprove(input)

// 	want := "FizzBuzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn16WhenInput16(t *testing.T) {
// 	input := 16

// 	got := FizzBuzzImprove(input)

// 	want := "16"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn17WhenInput17(t *testing.T) {
// 	input := 17

// 	got := FizzBuzzImprove(input)

// 	want := "17"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn18WhenInput18(t *testing.T) {
// 	input := 18

// 	got := FizzBuzzImprove(input)

// 	want := "Fizz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturn20WhenInput20(t *testing.T) {
// 	input := 20

// 	got := FizzBuzzImprove(input)

// 	want := "Buzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }

// func TestFizzBuzzShouldReturnFizzBuzzWhenInput30(t *testing.T) {
// 	input := 30

// 	got := FizzBuzzImprove(input)

// 	want := "FizzBuzz"
// 	if got != want {
// 		t.Errorf("got %q but want %q", got, want)
// 	}
// }
