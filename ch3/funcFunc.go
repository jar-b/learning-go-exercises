package main

import "fmt"

/*
1. Write a function that returns a function that performs a +2 on integers. Name the function plusTwo. You should then be able do the following:

p := plusTwo()
fmt.Printf("%v\n", p(2))

2. Generalize the function from above and create a plusX(x) which returns functions that add x to an integer.
*/

func plusTwo() (p func(int) int) {
    return func(n int) (n2 int) {
        return (n + 2)
    }
}

func plusX(x int) (p func(int) int) {
    return func(n int) (nx int) {
        return (n + x)
    }
}

func main() {
    p := plusTwo()
    fmt.Printf("%v\n", p(2))

    x := plusX(5)
    fmt.Printf("%v\n", x(5))
    fmt.Printf("%v\n", x(8))
}
