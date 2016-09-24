package main
import "fmt"

func findGreatest(nums ...int) int {
  greatest := 0

  for _, val := range nums {
    if greatest < val {
      greatest = val
    }
  }

  return greatest
}



func main() {
  fmt.Println(findGreatest(12,5432,123,654,231,6,4))
}