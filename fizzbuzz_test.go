package fizzbuzz

import "testing"

func TestFizzBuzzOneSholdGetOne(t *testing.T) {
	result := FizzBuzz(1)
	expected := "1"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetTwo(t *testing.T) {
	result := FizzBuzz(2)
	expected := "2"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetThree(t *testing.T) {
	result := FizzBuzz(3)
	expected := "Fizz"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetFour(t *testing.T) {
	result := FizzBuzz(4)
	expected := "4"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetFive(t *testing.T) {
	result := FizzBuzz(5)
	expected := "Buzz"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetSix(t *testing.T) {
	result := FizzBuzz(6)
	expected := "Fizz"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetSeven(t *testing.T) {
	result := FizzBuzz(7)
	expected := "7"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetEight(t *testing.T) {
	result := FizzBuzz(8)
	expected := "8"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}

func TestFizzBuzzOneSholdGetNine(t *testing.T) {
	result := FizzBuzz(9)
	expected := "Fizz"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
