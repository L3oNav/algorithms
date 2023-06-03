package main

import "fmt"

func main() {
    nums := [5]int{10, 20, 30, 40, 50}

    for index, value := range nums {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    length := len(nums)
    fmt.Println("Length of array:", length)

    array_copy := make([]int, len(nums))
    copy(array_copy, nums[:])
    fmt.Println("Copy of array:", array_copy)


    trimmed := nums[1:4]
    fmt.Println("Array recortado:", trimmed)

    nums[2] = 100
    fmt.Println("Array modificado:", nums)
}

