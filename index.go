package main

import "fmt"

// func main() {
// 	fmt.Println("hello world")

// }
func main() {
	
// var student1 string = "John" //type is string
//   var student2 = "Jane" //type is inferred
//   x := 2 //type is inferred

//   fmt.Println(student1)
//   fmt.Println(student2)
//   fmt.Println(x)
var a string = "hello world" //type is inferred
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  var i string = "Hello"
  var j int = 15

  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}