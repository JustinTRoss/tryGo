package main
import "fmt"

func swap(numA *int, numB *int) {
  temp := *numB
  *numB = *numA
  *numA = temp
}

func main() {
  x := 12
  y := 25
  fmt.Println('x', x, 'y', y)
  swap(&x, &y)
  fmt.Println('x', x, 'y', y)
}