package Array

import (
    "fmt"
)

func Example() {
    // Update the FixedArray variable with 5 strings
    UpdateArray("apple", "banana", "cherry", "date", "elderberry")

    // Print the current state of the FixedArray
    fmt.Printf("[Array]: %v \n", GetArray()) // Output: [apple banana cherry date elderberry]

    // Update the FixedArray with 5 strings
    UpdateArray("fig", "grape", "kiwi", "lemon", "mango")

    // Print the updated state of the FixedArray
    fmt.Printf("[Array]: %v \n\n", GetArray()) // Output: [fig grape kiwi lemon mango]
}