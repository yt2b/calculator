package main

import (
	"bufio"
	"calculator/parser"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		scanner.Scan()
		formula := scanner.Text()
		if formula == "" {
			return
		}
		value, err := parser.NewParser(formula).Parse()
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		fmt.Printf("%v\n", value)
	}
}
