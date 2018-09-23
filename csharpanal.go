package main

import (
	"io/ioutil"
	"fmt"
	)

func main() {
	input, err := ioutil.ReadFile("in.txt")
	if err != nil {
		fmt.Printf("could not read input file: %v\n", err)
		return
	}
	input = DeleteComments(input)

	output := ""
	l := new(Lexer)
	l.LoadText(input)
	for l.NextToken() {
		token, err := l.GetToken()
		if err != nil {
			fmt.Printf("could not get token: %v\n", err)
			return
		}
		if l.tipe == PUNCTUATOR {
			output += "PUCTUATOR  : "
		} else if l.tipe == OPERATOR {
			output += "OPERATOR   : "
		} else if l.tipe == LITERAL {
			output += "LITERAL    : "
		} else if l.tipe == KEYWORD {
			output += "KEYWORD    : "
		} else if l.tipe == IDENTIFIER {
			output += "IDENTIFIER : "
		}
		output += token + "\n"
	}

	ioutil.WriteFile("out.txt", []byte(output), 777)

	return
}