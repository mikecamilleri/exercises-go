package main

import "testing"

func TestFizzerBuzzerDivisibleByThree(t *testing.T) {
	var s string
	s = fizzerBuzzer(3)
	if s != "Fizz" {
		t.Errorf("fizzerBuzzer(3) returned: \"%s\", want: \"Fizz\".", s)
	}
	s = fizzerBuzzer(6)
	if s != "Fizz" {
		t.Errorf("fizzerBuzzer(6) returned: \"%s\", want: \"Fizz\".", s)
	}
}

func TestFizzerBuzzerDivisibleByFive(t *testing.T) {
	var s string
	s = fizzerBuzzer(-5)
	if s != "Buzz" {
		t.Errorf("fizzerBuzzer(-5) returned: \"%s\", want: \"Buzz\".", s)
	}
	s = fizzerBuzzer(5)
	if s != "Buzz" {
		t.Errorf("fizzerBuzzer(5) returned: \"%s\", want: \"Buzz\".", s)
	}
	s = fizzerBuzzer(10)
	if s != "Buzz" {
		t.Errorf("fizzerBuzzer(10) returned: \"%s\", want: \"Buzz\".", s)
	}
}

func TestFizzerBuzzerDivisibleByThreeAndFive(t *testing.T) {
	var s string

	s = fizzerBuzzer(-15)
	if s != "FizzBuzz" {
		t.Errorf("fizzerBuzzer(-15) returned: \"%s\", want: \"FizzBuzz\".", s)
	}
	s = fizzerBuzzer(0) // zero is divisible by any number
	if s != "FizzBuzz" {
		t.Errorf("fizzerBuzzer(0) returned: \"%s\", want: \"FizzBuzz\".", s)
	}
	s = fizzerBuzzer(15)
	if s != "FizzBuzz" {
		t.Errorf("fizzerBuzzer(15) returned: \"%s\", want: \"FizzBuzz\".", s)
	}
	s = fizzerBuzzer(30)
	if s != "FizzBuzz" {
		t.Errorf("fizzerBuzzer(30) returned: \"%s\", want: \"FizzBuzz\".", s)
	}
}

func TestFizzerBuzzerNotDivisibleByThreeOrFive(t *testing.T) {
	var s string
	s = fizzerBuzzer(-22)
	if s != "" {
		t.Errorf("fizzerBuzzer(-22) returned: \"%s\", want: \"\".", s)
	}
	s = fizzerBuzzer(22)
	if s != "" {
		t.Errorf("fizzerBuzzer(22) returned: \"%s\", want: \"\".", s)
	}
}
