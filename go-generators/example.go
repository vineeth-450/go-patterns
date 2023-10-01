/*
Sample implementation of generators similar to Python in Golang
Approach :	Use the concept of go routines to make computations to generate.
		Pass a channel to the generator function and put value into channel to "yield" a value from function.
		Spawn a routine for the generator function and range over the channel passed to the generator function to read generated values.

Run this program: https://play.golang.org/p/zr2wfpPN291
*/

package main

import (
	"fmt"
)

// Generator function. Pass a channel to the generator function and put value into channel to "yield" a value from function.
func squaresGenerator(c chan int) {
	for i := 1; ; i++ {
		c <- i * i
	}
}

func main() {
	c := make(chan int)

	// Spawn a routine for the generator function.
	go squaresGenerator(c)

	// Range over the channel passed to the generator function to read generated values.
	for sqr := range c {
		if sqr > 100 {
			break
		}
		fmt.Println("Next Square is", sqr)
	}
}
