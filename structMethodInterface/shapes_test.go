package main

import (
    "testing"
    "math"
)

func TestPerimeter(t *testing.T) {
    rect := Rectangle{10.0, 10.0}
    got := rect.Perimeter()
    want := 40.0

    if got != want {
        t.Errorf("got %.2f but expected %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    
    areaTest := []struct {
        name string
        shape Shape
        hasArea float64
    }{
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 10}, hasArea: 120.0},
        {name: "Circle", shape: Circle{Radius: 5}, hasArea: 25 * math.Pi},
        {name: "Triangle", shape: Triangle{Base: 6, Height: 8}, hasArea: 24.0},
    }

    for _, tt := range areaTest{
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %g but expected %g", tt.name, got, tt.hasArea)
            }
        })
    }
}
