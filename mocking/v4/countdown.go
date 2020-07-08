package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

//Sleeper is an interface
type Sleeper interface {
	Sleep()
}

//SpySleeper in a struct & implement Sleeper
type SpySleeper struct {
	Calls int
}

//Sleep means thread sleep times
func (s *SpySleeper) Sleep() {
	s.Calls++
}

//ConfigurableSleeper is a struct
type ConfigurableSleeper struct {
	duration time.Duration
}

//Sleep for time by given time
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

//CountdownOperationsSpy is a struct
type CountdownOperationsSpy struct {
	Calls []string
}

//Sleep add record to Spy
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

//CountDown is a timer "3 2 1 Go!"
func CountDown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	CountDown(os.Stdout, sleeper)
}
