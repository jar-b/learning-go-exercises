package main

import "fmt"

/*
Q9. (1) Stack
    Create a simple stack which can hold a fixed number of ints. It does not have to
    grow beyond this limit. Define push – put something on the stack – and pop –
    retrieve something from the stack – functions. The stack should be a LIFO (last
    in, first out) stack.
*/

type stack struct {
    i int
    d [10]int
}

func (s *stack) push(v int) {
    // add item to top of stack
    s.d[s.i] = v
    s.i++
}

func (s *stack) pop() (v int) {
    // pull item from top of stack
    s.i--
    v = s.d[s.i]
    // empty last element
    s.d[s.i] = 0

    return v
}

func main() {
    var s stack
    fmt.Println("Empty stack: ", s)

    // push
    s.push(5)
    s.push(12)
    s.push(65)
    fmt.Println("After push: ", s)

    // pop 
    s.pop()
    fmt.Println("After pop: ", s)
}
