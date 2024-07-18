package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {


    assertSum := func(t testing.TB, got, want int, numbers []int) {
        t.Helper()
        if got != want {
            t.Errorf("got %d but want %d, given %v", got, want, numbers)
        }
    }

    t.Run("collection of 5 numbers", func(t *testing.T) {
        numbers := []int{1,2,3,4,5}

        got := Sum(numbers)
        want := 15
        
        assertSum(t, got, want, numbers)
    })
    t.Run("collection of any size", func(t *testing.T) {
        numbers := []int{1, 2, 3}

        got := Sum(numbers)
        want := 6

        assertSum(t, got, want, numbers)
    })
}

func TestSumAll(t *testing.T) {
    
    got := SumAll([]int{1,2}, []int{0,9})
    want := []int{3,9}

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v but wanted %v", got, want)
    }
}

func TestSumTails(t *testing.T) {

    checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v but expected %v", got, want)
        }   
    }

    t.Run("make the sums of some slices", func(t *testing.T) {
        got := SumTails([]int{3, 2}, []int{4, 5})
        want := []int{2, 5}
        
        checkSums(t, got, want)
    })
    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumTails([]int{}, []int{3,4 ,5})
        want := []int{0,9}

        checkSums(t, got, want)
    })
}

