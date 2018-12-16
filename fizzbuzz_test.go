package fizzbuzz

import "testing"

func TestFizzBuzzOneSholdGetOne(t *testing.T) {
	result := FizzBuzz(1)
	expected := "1"
	if result != expected {
		t.Errorf("it should return %s but get %s", expected, result)
	}
}
