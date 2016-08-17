package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const maxIterations = 30000

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	fmt.Println()
	interpret("++++++++[>++++++++<]>+.")
}

func interpret(bfCodeString string) {
	bfCode := []rune(bfCodeString)
	var cells [30000]byte
	cellptr := 0
	pstack := NewPtrStack()
	var input []byte
	//iter := 0
	for i := 0; i < len(bfCode); i++ {
		switch bfCode[i] {
		case '<':
			cellptr--
			cellptr %= len(cells)
		case '>':
			cellptr++
			cellptr %= len(cells)
		case '+':
			cells[cellptr]++
		case '-':
			cells[cellptr]--
		case '[':
			if cells[cellptr] == 0 {
				openBrak := 0
				i++
				for ; !(openBrak == 0 && bfCode[i] == ']'); i++ {
					if bfCode[i] == '[' {
						openBrak++
					}
					if bfCode[i] == ']' {
						openBrak--
					}
				}
				//iter = 0
				continue
			}
			pstack.Push(i)
			//iter++
		case ']':
			/*if iter == maxIterations {
				log.Panic("ERROR: TOO MANY ITERATIONS")
			}*/
			if !pstack.IsEmpty() {
				i = pstack.Pop() - 1
				fmt.Println(len(cells), ";", cellptr)
				time.Sleep(1)
			}
		case '.':
			fmt.Print(string(cells[cellptr]))
		case ',':
			if len(input) == 0 {
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				input = []byte(scanner.Text())
			}
			if len(input) != 0 {
				cells[cellptr] = input[0]
				input = input[1:]
			}
		}
	}
}
