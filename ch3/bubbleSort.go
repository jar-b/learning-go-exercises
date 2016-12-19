package main

import "fmt"

/*

Write a function that performs a bubble sort on a slice of ints. From :

    It works by repeatedly stepping through the list to be sorted, comparing each pair of adjacent items and swapping them if they are in the wrong order. The pass through the list is repeated until no swaps are needed, which indicates that the list is sorted. The algorithm gets its name from the way smaller elements "bubble" to the top of the list.

*/


func bubble(nums []int) (sorted []int) {
    // initialize a new array, then copy original contents (this preserves the unordered original)
    sorted = make([]int, len(nums))
    copy(sorted, nums)

    swapped := true
    num_iter := 0

    for swapped {
        // on start of a new loop, initialize swapped as false
        swapped = false
        num_iter++

        for i := 0; i < (len(sorted) - 1); i++ {
            // if next number is smaller than current
            if sorted[(i+1)] < sorted[i] {
                // swap them
                sorted[(i+1)], sorted[i] = sorted[i], sorted[(i+1)]
                // indicate a swap occurred
                swapped = true
            }
        }
    }

    fmt.Printf("Finished sorting after %d iterations.\n", num_iter)

    return sorted
}

func main() {
    // Test1
    nums := []int{6, 12, 4, 3, 7}
    sorted := bubble(nums)
    fmt.Println("Original: ", nums)
    fmt.Println("Sorted: ", sorted)

    // Test2
    nums2 := []int{51, 32, 111, 65, 14, 122, 78, 412, 83, 6, 182, 19, 202}
    sorted2 := bubble(nums2)
    fmt.Println("Original: ", nums2)
    fmt.Println("Sorted: ", sorted2)
}
