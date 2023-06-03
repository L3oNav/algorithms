package main

import "fmt"

func main() {
    // Declaration of a fixed length array
    var nums [5]int

    // Assigning values to array elements
    nums[0] = 10
    nums[1] = 20
    nums[2] = 30
    nums[3] = 40
    nums[4] = 50

    fmt.Println(nums[0])  // 10
    fmt.Println(nums[2])  // 30

    for i := 0; i < len(nums); i++ {
        fmt.Println(nums[i])
    }
}

