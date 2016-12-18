package main

import "fmt"

/*
Write a function that finds the maximum value in an int slice ([]int).
*/

func maxVal(vals []int) (max int) {
    // initialize max as the first value
    max = vals[0]

    for _, v := range vals {
        if v > max {
            max = v
        }
    }

    return max
}

func main() {
    vals := []int{1, 5, 12, 6, 2}
    max := maxVal(vals)

    fmt.Printf("Max value is: %d\n", max)
}
