package main

import "fmt"

type Node struct {
    value int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (list *LinkedList) Insert(value int) {
    newNode := &Node{value: value}
    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func (list *LinkedList) Print() {
    current := list.head
    for current != nil {
        fmt.Println(current.value)
        current = current.next
    }
}

func main() {
    list := LinkedList{}
    list.Insert(10)
    list.Insert(20)
    list.Insert(30)
    list.Print()
}
