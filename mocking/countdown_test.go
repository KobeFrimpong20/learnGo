package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown (t *testing.T) {
    t.Run("Sleep is called 3 times and output is correct", func (t *testing.T) {
        buffer := &bytes.Buffer{}
        spySleeper := &SpyCountdownOperations{}

        Countdown(buffer, spySleeper)
       
        got := buffer.String()
        want := `3
2
1
Go!`

        if got != want {
           t.Errorf("Expected %q, but received %q as output", want, got) 
        }
    })
    t.Run("Output and sleep operations happen in correct order", func (t *testing.T) {
        spyCountdown := SpyCountdownOperations{}
        Countdown(&spyCountdown, &spyCountdown)

        want := []string{
            write,
            sleep, 
            write, 
            sleep, 
            write,
            sleep, 
            write,
        }

        if !reflect.DeepEqual(want, spyCountdown.Calls) {
            t.Errorf("wanted calls %v, but got %v", want, spyCountdown.Calls)
        }
    })
}

func TestConfigurableSleeper(t *testing.T) {
    sleepTime := 5 * time.Second

    spyTime := &SpyTime{}
    sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
    sleeper.Sleep()

    if spyTime.durationSlept != sleepTime {
        t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
    }
}
