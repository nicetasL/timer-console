package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	log.Printf("Timer started for %d seconds...\n", t.duration)
	time.Sleep(time.Duration(t.duration) * time.Second)
	log.Printf("Timer of %d seconds finished!\n", t.duration)
}

func main() {
	// Parse command-line arguments
	duration := flag.Int("time", 0, "Duration of the timer in seconds")
	flag.Parse()

	// Check if duration is valid
	if *duration <= 0 {
		fmt.Println("Error: Please provide a positive duration in seconds.")
		os.Exit(1)
	}

	// Create and start the timer
	timer := NewTimer(*duration)
	timer.Start()
}
