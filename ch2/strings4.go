package main

import "fmt"

/*
4. Write a Go program that reverses a string, so “foobar” is printed as “raboof”. Hint:
You will need to know about conversion; skip ahead to section “Conversions” on
page 54.
*/

func reverseString(s string) string {
    // convert string to byte representation
    b := []byte(s)

    // reverse the order
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }

    // convert back to string and return
    rev := string(b)
    return rev
}

func main() {
    var s string = "foobar"
    rev := reverseString(s)
    fmt.Println(rev)
}
