package main

import (
	"bufio"
	"fmt"
	"os"

	"simpleInterpreter/Interpreter"
	"simpleInterpreter/Lexer"
	"simpleInterpreter/Parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">>> ")
		line, err := reader.ReadBytes('\n')
		text := string(line)[:len(line)-1]

		if err != nil {
			fmt.Println(err)
			continue
		}

		lexer := Lexer.NewLexer(text)
		parser, err := Parser.NewParser(lexer)

		if err != nil {
			fmt.Println(err)
			continue
		}

		interpreter := Interpreter.NewInterpreter(parser)
		result, err := interpreter.Interpret()

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)

	}
}
