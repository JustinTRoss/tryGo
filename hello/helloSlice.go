package main

import "fmt"

func main() {
  // Create slice that contains [1,2,3]
  slice1 := []int{1,2,3}
  // Create slice that contains slice1 with [4,5] appended
  slice2 := append(slice1, 4, 5, 6, 7, 8, 9)
  // Create slice with max size of 2
  slice3 := make([]int, 2, 6)
  // Copy slice2 onto slice3. This results in [1,2] because of max.
  copy(slice3, slice2[3:])
  fmt.Println(slice1, slice2, slice3)
  copy(slice1, slice3)
  fmt.Println(slice1, slice2, slice3)

}