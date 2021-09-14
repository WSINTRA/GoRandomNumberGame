// First attempt with Go, Sep 2021 about to start a new job and want to gain GoLang experience
// Random number guessing game based on the HEAD first Go - A brain friendly guide book
// Chapter 2
// Improvements include:
// Set guesses as 10 instead of 0, so guess can be increased more easily by changing the variable

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var guesses = 10
	var complete = false
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1

	fmt.Println("Guess the random number between 1-100")
	for i := 0; i < guesses; i++ {
		fmt.Printf("You have %v guess left \n", (guesses - i))

		if (guesses - i) == 1 {
			fmt.Println("Last chance!!")
		} else {
			fmt.Println("type a number:")
		}
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("too low")
		} else if guess > target {
			fmt.Println("Too high")
		} else {
			fmt.Println("Magic, how you guess that quickly!")
			complete = true
			break
		}
	}

	if !complete {
		fmt.Printf("Sorry the correct number was %v \n", target)
	}
}
