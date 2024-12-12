package main

import (
	"fmt"
)

func cook(dish string, channel chan string) {
	
	fmt.Printf("Start Cooking %s \n", dish)
	// cooking something
	fmt.Printf("End Cooking %s \n", dish)

	channel <- "I'm Cooking " + dish;
}
func main() {
	fmt.Println("Starting Main Function")
	channel := make(chan string)
	go cook("pasta", channel);
	go cook("pizza",channel)
	fmt.Println(<- channel)
	fmt.Println("Completing Main Function")
}

// ? Parallelism and Concurrency.
