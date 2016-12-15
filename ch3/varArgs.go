package main

import "fmt"


/* Write a function that takes a variable number of ints and print each integer on a separate line. */


func varArgs(nums ...int) {
    for _, n := range nums {
        fmt.Println(n)
    }
}


func main() {
    varArgs(1,2,3)
    varArgs(22,1234,23,1235)
}
