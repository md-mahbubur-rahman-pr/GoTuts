package main

import "fmt"

func main() {

    var s = "this is a string"
    fmt.Println(s)

    var p, q, r int = 1, 2, 3
    fmt.Println(p, q, r)

    var flag = true
    fmt.Println(flag)

    var i int
    fmt.Println(i)
    
    var f float32
    fmt.Println(f)

    // The := syntax is shorthand for declaring and initializing a variable, e.g. for var str string = "this is another string" in this case.
    str := "this is another string"
    fmt.Println(str)
}
