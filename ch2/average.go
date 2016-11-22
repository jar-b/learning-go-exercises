package main

import "fmt"

/*
Write code to calculate the average of a float64 slice. In a later exercise (Q6 you
will make it into a function.
*/

func calcAverage(data []float64) float64 {
    sum := float64(0)

    for _, num := range data {
        sum += num
    }

    avg := sum / float64(len(data))
    return avg
}

func main() {
    data := []float64{12, 24, 36}
    avg := calcAverage(data)
    fmt.Println("Data: ", data)
    fmt.Println("Average: ", avg)
}
