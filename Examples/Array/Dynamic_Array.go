package Array

import "fmt"

// Dynamic_Array_Example demonstrates the usage of a dynamic array in Go.
// The function creates a slice with an initial capacity of 5, then appends
// values to the slice, prints the original and modified slice, and iterates
// over the slice to print its elements.
func Dynamic_Array_Example() {
    // A slice with an initial capacity of 5
    var slice []string

    // Appending values to the slice
    slice = append(slice, "h")
    slice = append(slice, "e")
    slice = append(slice, "l")
    slice = append(slice, "l")
    slice = append(slice, "o")

    // Printing the slice
    fmt.Println("Slice:", slice)

    // Modifying an element of the slice
    slice[3] = "p"

    // Printing the modified slice
    fmt.Println("Modified slice:", slice)

    // Iterating over the slice
    fmt.Print("Slice elements: ")
    for _, element := range slice {
        fmt.Printf("%s ", element)
    }
}
