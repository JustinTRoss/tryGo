package main
import "fmt"

func makeEvenGenerator() func() int {
  i := 0
  return func() int {
    i += 2

    return i
  }
}

func main() {
  evensOne := makeEvenGenerator()
  evensTwo := makeEvenGenerator()
  fmt.Println(evensOne())
  fmt.Println(evensOne())
  fmt.Println(evensOne())
  fmt.Println(evensTwo())
  fmt.Println(evensTwo())
  fmt.Println(evensOne())
  fmt.Println(evensTwo())
}