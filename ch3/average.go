package main

import "fmt"

/*
Q6. Write a function that calculates the average of a float64 slice
*/

func average(s []float64) (avg float64) {
    var total float64
    for _, v := range s {
        total += v
    }

    return total / float64(len(s))
}

func main() {
    test := []float64{12, 24, 36}
    avg := average(test)
    fmt.Printf("Slice is: %v\n", test)
    fmt.Printf("Average is: %v\n", avg)
}

