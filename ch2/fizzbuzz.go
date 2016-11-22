package main

import "fmt"

func main() {
    // Using if/elif/else
    fmt.Println("\nUsing if/else...")
    for i := 1; i <= 100; i++ {
        if (i % 3 == 0) && (i % 5 == 0) {
            fmt.Println("fizzbuzz")
        } else if (i % 3 == 0)  {
            fmt.Println("fizz")
        } else if (i % 5 == 0)  {
            fmt.Println("buzz")
        } else {
            fmt.Println(i)
        }
    }

    // Using switch
    fmt.Println("\nUsing switch...")
    for i := 1; i <= 100; i++ {
        switch {
            case (i % 3 == 0) && (i % 5 == 0):
                fmt.Println("fizzbuzz")
            case (i % 3 == 0):
                fmt.Println("fizz")
            case (i % 5 == 0):
                fmt.Println("buzz")
            default:
                fmt.Println(i)
        }
    }
}
