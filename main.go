package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("programTest.txt")
	if err != nil {
		panic(err)
	}
	program := NewLexer(file)

	var test []Result
	for {

		pos, token, lit := program.Analyzer()

		if token == EOF {
			break
		}
		var result Result
		result.lit = lit
		result.token = token
		result.pos = pos
		test = append(test, result)
	}
	fmt.Println("-------------------------------------------------------------------------------------")
	for i := 0; i < len(test); i++ {

		fmt.Println(test[i])

		fmt.Println("-----------------------------------------------------------------------------------")
	}
	fmt.Println("Precione qualquer tecla para finalizar!")
	fmt.Scanln()
}
