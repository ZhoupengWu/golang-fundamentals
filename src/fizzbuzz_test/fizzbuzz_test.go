package fizzbuzztest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzbuzz (t *testing.T) {
	t.Run("When1_Then1", func(t *testing.T) {
		res := Fizzbuzz(1)
		assert.Equal(t, "1", res)
	})

	t.Run("When3_ThenFizz", func(t *testing.T) {
		res := Fizzbuzz(3)
		assert.Equal(t, "Fizz", res)
	})

	t.Run("When5_ThenBuzz", func(t *testing.T) {
		res := Fizzbuzz(5)
		assert.Equal(t, "Buzz", res)
	})

	t.Run("When15_ThenFizzBuzz", func(t *testing.T) {
		res := Fizzbuzz(15)
		assert.Equal(t, "FizzBuzz", res)
	})
}

/* func TestFizzbuzz_1 (t *testing.T) {
	res := Fizzbuzz(1)
	assert.Equal(t, "1", res)
} */

/* func TestFizzbuzz_3 (t *testing.T) {
	res := Fizzbuzz(3)
	assert.Equal(t, "Fizz", res)
} */

/* func TestFizzbuzz (t *testing.T) {
	var input int = 1

	res := Fizzbuzz(input)

	if res != "1" {
		t.Errorf("Expected %q, got %s", input, res)
	}
} */