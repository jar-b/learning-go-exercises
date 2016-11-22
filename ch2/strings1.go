package main

import (
    "fmt"
    "bytes"
)


func main() {
    fmt.Println("Question 1...")

    // Using simple string append
    var s string
    for i := 0; i < 10; i++ {
        s += "A"
        fmt.Println(s)
    }

    // Using bytes.Buffer
    var b bytes.Buffer
    for i := 0; i < 10; i++ {

        b.Write([]byte("A"))
        fmt.Println(b.String())

        b.WriteString("A")
        fmt.Println(b.String())
    }
}
