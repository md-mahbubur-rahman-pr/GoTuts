package main 

import "fmt"

// const declares a constant value
const N 9999
const F 1.255

func main(){
  fmt.Println(N);
  
  // A const statement can appear anywhere like var statement
  const str String = "This is a constant"
  fmt.Println(str)
}
