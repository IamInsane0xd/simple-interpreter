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
	if len(os.Args) <= 1 {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print(">>> ")
			line, err := reader.ReadBytes('\n')

			if err != nil {
				fmt.Println(err)
				continue
			}

			text := string(line)[:len(line)-1]
			result, err := runInterpreter(text)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(result)
		}
	} else {
		fileName := os.Args[1]
		buf, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Println(err)
			return
		}

		text := string(buf)
		result, err := runInterpreter(text)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(result)
	}
}

func runInterpreter(text string) (int64, error) {
	lexer := Lexer.NewLexer(text)
	parser, err := Parser.NewParser(lexer)

	if err != nil {
		return 0, err
	}

	interpreter := Interpreter.NewInterpreter(parser)
	result, err := interpreter.Interpret()

	if err != nil {
		return 0, err
	}

	return result, nil
}
