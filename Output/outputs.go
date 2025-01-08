package main

import "fmt"

const a = 5     // constant declaration
const b = 10    // constant declaration
const c = a + b // constant declaration

func main() {
	fmt.Print(a)                                    // print without newline
	fmt.Print("\n")                                 // print newline
	fmt.Println(b)                                  // print with newline
	fmt.Printf("The sum is: %v and Type: %T", c, c) // formatted print
}
