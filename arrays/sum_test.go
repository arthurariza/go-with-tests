package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sums all numbers in one slice into a new slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumAll(numbers)
		want := []int{6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v given: %v", got, want, numbers)
		}
	})

	t.Run("sums all numbers in all slices into a new slice", func(t *testing.T) {
		numbersOne := []int{1, 2, 3}
		numbersTwo := []int{1, 1, 1}

		got := SumAll(numbersOne, numbersTwo)
		want := []int{6, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	})
}

func TestSumTail(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := SumTail(numbers)
	want := 14

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sums all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("sums all tails with different sizes", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 1})
		want := []int{5, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
