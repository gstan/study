package main

import "fmt"

func main() {
    input := []int{9,3,5,4,2,7,8,1,6}
    fmt.Println(input)
    output := selectSort(input)
    fmt.Println(output)
}

func selectSort(input []int) []int {
    length := len(input)
    if length  <= 0 {
        return input
    }
    for i := 0; i < length-1; i++ {
        min := i
        for j := i + 1; j < length; j++ {
            if input[min] > input[j] {
                min = j
            }
        }
        input[i], input[min] = input[min], input[i]
    }
    return input
}
