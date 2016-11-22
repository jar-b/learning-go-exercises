package main

import "fmt"

/*
Q7. (0) Integer ordering
    Write a function that returns its (two) parameters in the right, numerical (ascend-
    ing) order:
    f(7,2) → 2,7
    f(2,7) → 2,7
*/

func order(n1, n2 int) (r1, r2 int) {
    if n1 <= n2 {
        return n1, n2
    } else {
        return n2, n1
    }
}

func main() {
    fmt.Println(order(7,2))
    fmt.Println(order(2,7))
}
