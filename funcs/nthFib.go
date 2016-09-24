package main
import "fmt"

func nthFib(num uint) uint {
  ary := [2]uint{0,1}
  largestInd := 0

  if num == 0 || num == 1 {
    return ary[num]
  }

  for i := 2; uint(i) <= num; i++ {  
    if i % 2 == 0 {
      largestInd = 0
      ary[0] = ary[0] + ary[1]
    } else {
      largestInd = 1
      ary[1] = ary[0] + ary[1]
    }
  }
  return ary[largestInd]
}

func main() {
  fmt.Println(nthFib(2000000000))
}