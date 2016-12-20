package stack

import "testing"

func TestEmptyStack(t *testing.T) {
    var s Stack

    if s.i != 0 {
        t.Log("Empty stack should have i = 0!")
        t.Fail()
    }
}

func TestPush(t *testing.T) {
    var s Stack
    s.Push(12)

    if s.i != 1 {
        t.Log("Index (i) after one Push should be 1!")
        t.Fail()
    }
    if s.d[0] != 12 {
        t.Log("First data element (d[0]) after one Push should be 12!")
        t.Fail()
    }
}

func TestPushPop(t *testing.T) {
    var s Stack
    s.Push(12)
    s.Push(61)
    s.Push(32)

    if s.i != 3 {
        t.Log("Index (i) after three Push's should be 3!")
        t.Fail()
    }

    s.Pop()

    if s.i != 2 {
        t.Log("Index (i) after three Push's and one Pop should be 2!")
        t.Fail()
    }
    if (s.d[0] != 12 && s.d[1] != 61) {
        t.Log("Incorrect data (d) after three Push's and one Pop!")
        t.Fail()
    }
}

