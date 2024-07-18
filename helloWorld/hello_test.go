package main

import (
    "testing"
)

func TestHello(t *testing.T) {
    t.Run("Saying hello with name", func(t *testing.T) {
        got := hello("Kobe", "")
        want := "Hello, Kobe"

        assertCorrectMessage(t, got, want)
    })
    t.Run("empty string Saying hello", func(t *testing.T) {
        got := hello("", "")
        want := "Hello, World"

        assertCorrectMessage(t, got, want)
    })
    t.Run("in Spanish", func(t *testing.T) { 
        got := hello("Kobe", "Spanish")
        want := "Hola, Kobe"

        assertCorrectMessage(t, got, want)
    })
    t.Run("in French", func(t *testing.T) {
        got := hello("Kobe", "French")
        want := "Bonjour, Kobe"

        assertCorrectMessage(t, got, want)
    })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }   
}
