package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// Timer interface defines the Start method
type Timer interface {
	Start() // Starts the timer countdown
}

// SimpleTimer struct holds the duration of the timer
type SimpleTimer struct {
	duration int
}

// NewTimer creates a new instance of a Timer
func NewTimer(duration int) Timer {
	return &SimpleTimer{duration: duration}
}

// Start begins the timer and prints when it finishes
func (t *SimpleTimer) Start() {
	time.Sleep(time.Duration(t.duration) * time.Second)
	fmt.Printf("Timer of %d seconds finished!\n", t.duration)
}

func main() {
	// Parse command-line arguments
	duration := flag.Int("time", 0, "Duration of the timer in seconds")
	flag.Parse()

	if *duration <= 0 {
		log.Fatal("Please provide a positive duration in seconds.")
	}

	// Create and start the timer
	timer := NewTimer(*duration)
	timer.Start()
}
