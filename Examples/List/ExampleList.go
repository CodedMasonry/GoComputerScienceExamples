package linkedlist

import (
    "fmt"
)

func Example() {
    
    // Create a new linked list
    list := &LinkedList{}

    // Append some values to the list
    list.Append(1)
    list.Append(2)
    list.Append(3)

    // Prepend a value to the list
    list.Prepend(0)

    // Print the current state of the list
    current := list.head
    for current != nil {
        fmt.Printf("[LinkedList]: %v \n", current.value)
        current = current.next
    }
    // Output: 0 1 2 3
    // For spacing when Printing
    fmt.Printf("\n")
  
    // Delete a value from the list
    list.Delete(2)

    // Print the updated state of the list
    current = list.head
    for current != nil {
        fmt.Printf("[LinkedList]: %v \n", current.value)
        current = current.next
    }
    // Output: 0 1 3
}
