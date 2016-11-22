package main

import (
    "fmt"
    "unicode/utf8"
)

/*
2. Create a program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk
In addition, make it output the number of bytes in that string. Hint: Check out the utf8 package.

3. Extend/change the program from the previous question to replace the three runes
at position 4 with ’abc’.
*/


func main() {
    // Question 2
    fmt.Println("Question 2:")
    s := "asSASA ddd dsjkdsjs dk" //string representation
    s_byte := []byte(s) //byte representation
    s_rune := []rune(s) //rune representation

    // Count length of string, byte slice, and rune slice 
    fmt.Println(len(s))
    fmt.Println(len(s_byte))
    fmt.Println(len(s_rune))

    // Alternative method to count using the unicode/utf8 package 
    fmt.Println(utf8.RuneCountInString(s))


    // Question 3
    fmt.Println("Question 3:")
    fmt.Println("Before: ", string(s_byte))

    b := "abc"
    b_byte := []byte(b)
    copy(s_byte[3:], b_byte) // use copy to replace original with new characters
    fmt.Println("After: ", string(s_byte))
}
