package main

import "fmt"


func fibo(n int) (f []int) {
    // pre-populate the first two elements
    f = append(f, 1, 1)
    fmt.Printf("Running fibonacci sequence for %d iterations.\n", n)

    if n < 2{
        fmt.Println("Error: n must be at least 2!")
    } else {
        // compute fibonacci sequence for number of iterations specified
        for i := 1; i < (n-1); i++ {
            f = append(f, (f[i-1] + f[i]))
        }
    }

    return f
}

func main() {
    // this should error
    fibo(1)

    // these work
    fmt.Println(fibo(2))
    fmt.Println(fibo(10))
    fmt.Println(fibo(25))
    // the capacity of int "rolls over" at the end of this test, causing
    // values to become negative
    fmt.Println(fibo(100))
}
