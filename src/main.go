/*
Package main implements a regular expression parser.
*/
package main

import (
	"fmt"
	fsm "./fsm"
)

func main() {
	fsm := fsm.InitFSM(5)

	fsm.SetTransition(0,1,[]rune{'c'})
	fsm.SetTransition(1,1,[]rune{'a'})
	fsm.SetTransition(1,2,[]rune{'t'})

	fsm.SetTransition(0,3,[]rune{'d'})
	fsm.SetTransition(3,3,[]rune{'o'})
	fsm.SetTransition(3,4,[]rune{'g'})

	fsm.SetTerminals([]int{2, 4})

	fsm.PrintFSM()

    fmt.Println(fsm.ValidateWord("ct"))
	fmt.Println(fsm.ValidateWord("cat"))
	fmt.Println(fsm.ValidateWord("caaat"))
	fmt.Println(fsm.ValidateWord("ctt"))
	fmt.Println(fsm.ValidateWord("catt"))
	fmt.Println(fsm.ValidateWord("caaatt"))
	fmt.Println(fsm.ValidateWord("dg"))
	fmt.Println(fsm.ValidateWord("dog"))
	fmt.Println(fsm.ValidateWord("dooog"))
	fmt.Println(fsm.ValidateWord("dgg"))
	fmt.Println(fsm.ValidateWord("dogg"))
	fmt.Println(fsm.ValidateWord("dooogg"))
	fmt.Println(fsm.ValidateWord("elephant"))
	fmt.Println(fsm.ValidateWord("hamster"))
	fmt.Println(fsm.ValidateWord("cdaotg"))

	fmt.Println("Exiting...")
}
