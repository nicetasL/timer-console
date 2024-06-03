package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// SimpleTimer struct holds the duration of the timer
type SimpleTimer struct {
	duration int
}

// NewTimer creates a new timer instance
func NewTimer(duration int) *SimpleTimer {
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
	fmt.Printf("Timer started for %d seconds...\n", *duration)
	timer.Start()
}
