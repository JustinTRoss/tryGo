package main
import "fmt"

func add(num ...int) int {
  total := 0
  for _, val := range num {
    total += val
  }
  return total;
}

func main() {
  fmt.Println(add(11,12,31,41,51))
}