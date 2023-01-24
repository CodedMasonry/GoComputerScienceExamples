package linkedlist

import "fmt"

func Example() {
    // Initialize a new linked list
    list := &LinkedList{}

    // Append values to the linked list
    list.Append(1)
    list.Append(2)
    list.Append(3)
    list.Append(4)
    list.Append(5)

    // Print the linked list
    fmt.Println("Linked List:")
    current := list.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println()

    // Prepend a new value to the linked list
    list.Prepend(0)

    // Print the linked list
    fmt.Println("Linked List after Prepend:")
    current = list.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println()

    // Delete a value from the linked list
    list.Delete(3)

    // Print the linked list
    fmt.Println("Linked List after Delete:")
    current = list.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println()
}
