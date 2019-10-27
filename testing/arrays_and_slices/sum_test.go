package main

import (
	"reflect"
	"testing"
)

//  sum of slices
func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	want := 6

	if got != want {
		t.Errorf("got %d want %d ", got, want)
	}
}

// sum of multiple slices
func TestSumAll(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d ", got, want)
		}

	}

	t.Run("make the sums of slices", func(t *testing.T){
		got := SumAll([]int{1, 2}, []int{9, 9})
		want := []int{3, 18}

		checkSums(t, got, want)
	})

	t.Run("safely pass an empty slice", func(t *testing.T){
		got := SumAll([]int{}, []int{9, 9})
		want := []int{0, 18}

		checkSums(t, got, want)
	})

}


// test sum tails

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d ", got, want)
		}

	}

	t.Run("return the last number of the slice", func(t *testing.T){
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	
	})

	t.Run("safely return an emmpty slice", func(t *testing.T){
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	
	})
}
