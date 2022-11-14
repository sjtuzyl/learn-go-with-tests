package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countDownStart = 3
	write          = "write"
	sleep          = "sleep"
)

func Countdown(w io.Writer, spy Sleeper) {
	for i := countDownStart; i > 0; i-- {
		// fmt.Fprintln(w, i)
		// fmt.Fprintf(w, fmt.Sprint(i))
		// time.Sleep(1 * time.Second)
		spy.Sleep()
		fmt.Fprintln(w, i)
	}

	// for i := countDownStart; i > 0; i-- {
	// 	fmt.Fprintln(w, i)
	// }

	// time.Sleep(1 * time.Second)
	spy.Sleep()
	fmt.Fprintf(w, finalWord)
	return
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func main() {
	// spy := SpySleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
