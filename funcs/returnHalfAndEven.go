package main
import "fmt"

func halfAndEven(num int) (int, bool) {
  h := num/2
  e := h % 2 == 0
  return h,e
}

func main() {
  fmt.Println(halfAndEven(340))
  fmt.Println(halfAndEven(431))
  fmt.Println(halfAndEven(462))
  fmt.Println(halfAndEven(3567))
  fmt.Println(halfAndEven(434))
  fmt.Println(halfAndEven(57))
  fmt.Println(halfAndEven(1021))
}