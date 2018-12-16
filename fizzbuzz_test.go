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
