package integers

import (
    "testing" 
    "fmt"
)

func TestAdder(t *testing.T) {
    got := Add(5,3)
    want := 8
    
    assertSum(t, got, want)
}

func ExampleAdd() {
    fmt.Println(Add(4, 6))
    // Output: 10
}

func assertSum(t testing.TB, got, want int) {
    if got != want {
        t.Errorf("expected %d but got %d", want, got)
    }
}
