package main

import (
	"fmt"
	"time"
)

// task is a simple function that prints which task number is being executed.
func task(i int) {
	fmt.Println("Doing task", i)
}

func main() {
	// Launch 11 goroutines (from 0 to 10)
	// Each goroutine runs the 'task' function concurrently.
	for i := 0; i <= 10; i++ {
		go task(i) // 'go' keyword starts the function as a separate lightweight thread
	}

	// Sleep the main goroutine for 2 seconds
	// This gives enough time for all other goroutines to finish their execution
	// before the program exits.
	// Without this line, the main function may exit immediately,
	// stopping other goroutines before they run.
	time.Sleep(time.Second * 2)
}
