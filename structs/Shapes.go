package main
import ("fmt";"math")

type Shape interface {
  area() float64
  perimeter() float64
}

type Circle struct {
  Shape
  x, y, r float64
}

type Rectangle struct {
  Shape
  h, w float64
}

func main() {
  c := Circle{1,2,3}

}