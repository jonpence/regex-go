/*
Package main implements a regular expression engine.
*/
package main

import (
	"fmt"
)

func main() {

	p := initProgram()
	p.run()

	fmt.Println("Exiting...")
}
