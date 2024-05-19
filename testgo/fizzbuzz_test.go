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
