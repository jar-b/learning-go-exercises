package main

import "fmt"

/*
A map()-function is a function that takes a function and a list. The function is applied to each member in the list and a new list containing these calculated values is returned. Thus:

    map(f(),(a1,a2,…,an−1,an))=(f(a1),f(a2),…,f(an−1),f(an))

Write a simple map()-function in Go. It is sufficient for this function only to work for ints.
*/

// simple example function
func plusFour(i int) (p int) {
    return (i + 4)
}

func mapPlusFour(input []int) (output []int) {
    // initialize output array with the same size (and capacity, by default) as input
    output = make([]int, len(input))

    for i, v := range input {
        output[i] = plusFour(v)
    }

    return output
}

func main() {
    input := []int{1, 4, 6, 2}
    output := mapPlusFour(input)

    fmt.Println("Input: ", input)
    fmt.Println("Output: ", output)
}


