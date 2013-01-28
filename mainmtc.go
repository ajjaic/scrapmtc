package main

import "fmt"

func main() {
  b := newBus("01A")


  fmt.Println(b.routenum)
  fmt.Println(b.servtype)
  fmt.Println(b.jmin)
  for i, v := range b.stname {
    fmt.Println(i+1, v)
  }

}