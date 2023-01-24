package Array

import "fmt"

// fixedArray is a fixed length array of 5 strings.
var fixedArray [5]string

// Fixed_Array_Example demonstrates the usage of a fixed length array in Go.
// The function assigns values to the fixed array, prints the original and modified array, and iterates
// over the array to print its elements.
func Fixed_Array_Example() {
    // Assigning values to the fixed array
    fixedArray[0] = "h"
    fixedArray[1] = "e"
    fixedArray[2] = "l"
    fixedArray[3] = "l"
    fixedArray[4] = "o"

    // Printing the fixed array
    fmt.Println("Fixed array:", fixedArray)

    // Modifying an element of the fixed array
    fixedArray[3] = "p"

    // Printing the modified fixed array
    fmt.Println("Modified fixed array:", fixedArray)

    // Iterating over the fixed array
    fmt.Print("Fixed array elements: ")
    for _, element := range fixedArray {
        fmt.Printf("%s ", element)
    }
}
