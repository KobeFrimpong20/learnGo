package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"
const sleep = "sleep"
const write = "write"

type Sleeper interface {
    Sleep()
}

type SpyCountdownOperations struct {
    Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

type ConfigurableSleeper struct {
    duration time.Duration
    sleep func(time.Duration)
}

func (cSleeper *ConfigurableSleeper) Sleep () {
    cSleeper.sleep(cSleeper.duration)
}

type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
    s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        fmt.Fprintln(out, i)
        sleeper.Sleep()
    }
    fmt.Fprintf(out, finalWord)
}

func main () {
    sleeper := &ConfigurableSleeper{time.Second, time.Sleep}
    Countdown(os.Stdout, sleeper)
}
