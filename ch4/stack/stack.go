// A package implementing a simple stack. Up to 10 elements 
// can be held at a time, and the last element in is the
// first element out.
package stack


type Stack struct {
    i int
    d [10]int
}

// Push pushes a new element on the top of the stack
func (s *Stack) Push(v int) {
    s.d[s.i] = v
    s.i++
}

// Pop pulls the last item from the top of the stack
func (s *Stack) Pop() (v int) {
    s.i--
    v = s.d[s.i]

    // empty last element
    s.d[s.i] = 0

    return v
}
