package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/jonpence/regex-go/internal/regex"
)

type Program struct {
	regex     []Regex
	reader    *bufio.Reader
	running   bool
}

func initProgram() Program {
	return Program{[]Regex{}, bufio.NewReader(os.Stdin), true}
}

func (p Program) readInput() (string, bool) {
	input, err := p.reader.ReadString('\n')

	if err != nil {
		log.Println("Unable to read input, try again: ", err)
		return "", false
	}

	input = strings.TrimSuffix(input, "\n")

	return input, true
}

func (p Program) run() {
	for p.running {
		fmt.Println("(1) Add new regular expression.")
		fmt.Println("(2) List available regular expressions.")
		fmt.Println("(3) Validate string using a regular expression.")
		fmt.Println("(4) Quit.")

		choice, ok := p.readInput()

		if !ok {
			continue
		}

		switch choice {
			case "1": p.addRegex()
			case "2": p.listRegex()
			case "3": p.validateString()
			case "4": p.quit()
			default: log.Println(choice, " is not a valid option. Try again.")
		}
	}
}

func (p *Program) addRegex() {
	fmt.Println("Enter new regular expression:")

	expression, ok := p.readInput()

	if !ok {
		log.Println("Unable to create new regex. Try again.")
		return
	}

	newRegex, ok := newRegex(expression)

	if !ok {
		log.Println("Unable to create new regex. Try again.")
		return
	}

	p.regex = append(p.regex, newRegex)
}

func (p Program) listRegex() {
	for i, regex := range p.regex {
		fmt.Printf("%d. %s \n", i, regex.expression)
	}
}

func stringToInt(input string) int {
	num := 0

	for i := 0; i < len(input); i++ {
		num *= 10
		num += int(input[i] - 48)
	}

	return num
}

func (p Program) validateString() {
	fmt.Println("Enter number of regex you wish to use:")

	number, ok := p.readInput()

	if !ok {
		log.Println("Unable to validate string. Try again.")
		return
	}

	fmt.Println("Enter string you wish to validate:")

	input, ok := p.readInput()

	if !ok {
		log.Println("Unable to validate string. Try again.")
		return
	}

	if p.regex[stringToInt(number)].validate(input) {
		fmt.Println("Valid.")
	} else {
		fmt.Println("Invalid.")
	}
}

func (p *Program) quit() {
	p.running = false
}
