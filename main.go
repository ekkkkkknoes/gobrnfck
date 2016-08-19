package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const maxIterations = 30000

func main() {
	if len(os.Args) > 1 {
		if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			fmt.Println(err)
			return
		}
		bfCode, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		interpret(bfCode)
		return
	}
	fmt.Println("Error: Must specify brainf*ck file!")
	return
}

func interpret(bfCode []byte) {
	var cells [30000]byte
	cellptr := 0
	pstack := NewPtrStack()
	var input []byte
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
				continue
			}
			pstack.Push(i)
		case ']':
			if !pstack.IsEmpty() {
				i = pstack.Pop() - 1
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
