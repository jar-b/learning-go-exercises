package main

import "fmt"

func main() {
	// Part 1 - simple for loop printing the number of the loop it is on
	fmt.Println("Starting Part 1...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Part 2 - re-write the functionality above, but using "goto"
	fmt.Println("Starting Part 2...")
	i := 0
	startLoop:
		if i < 10 {
			fmt.Println(i)
			i++
			goto startLoop
		}

	// Part 3 - rewriting the loop to fill an array and print it each iteration
	fmt.Println("Starting Part 3...")
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)

}


