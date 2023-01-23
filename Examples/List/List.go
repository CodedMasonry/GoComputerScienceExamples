package linkedlist

// Node represents a single node in a linked list, with a value and a pointer to the next node.
type Node struct {
    value int
    next *Node
}

// LinkedList represents a linked list data structure, with a head and tail pointer.
type LinkedList struct {
    head *Node
    tail *Node
}

// Append adds a new node with the provided value to the end of the linked list.
func (l *LinkedList) Append(value int) {
    node := &Node{value: value}
    if l.tail == nil {
        l.head = node
        l.tail = node
    } else {
        l.tail.next = node
        l.tail = node
    }
}

// Prepend adds a new node with the provided value to the beginning of the linked list.
func (l *LinkedList) Prepend(value int) {
    node := &Node{value: value}
    if l.head == nil {
        l.head = node
        l.tail = node
    } else {
        node.next = l.head
        l.head = node
    }
}

// Delete removes the first node with the provided value from the linked list.
func (l *LinkedList) Delete(value int) {
    if l.head == nil {
        return
    }
    if l.head.value == value {
        l.head = l.head.next
        return
    }
    current := l.head
    for current.next != nil {
        if current.next.value == value {
            current.next = current.next.next
            break
        }
        current = current.next
    }
}
