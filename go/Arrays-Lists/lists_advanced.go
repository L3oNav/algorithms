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
    newNode := &Node{value:value}

    if list.head == nil {
        list.head = newNode
    } else {
        newNode.next = list.head
        list.head = newNode
    }
}

func (list *LinkedList) Remove(value int) {
    if list.head == nil {
       return
    }
    if list.head.value == value {
      list.head = list.head.next
      return
    }

    current := list.head
    for current.next != nil {
        if current.next.value == value {
            current.next = current.next.next
            return
        }
        current = current.next
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
    fmt.Println("Removing 20")
    list.Remove(20)
    fmt.Println("Removed 20")
    list.Print()
}
