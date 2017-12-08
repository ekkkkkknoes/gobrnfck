package main

import (
	"bufio"
	"fmt"
	"os"
)

//Interpreter is a brainfuck interpreter and should always be created with
//NewInterpreter(cellcount, maxiter int)
type Interpreter struct {
	cells         []byte
	maxIterations int
	cellptr       int
	pc            int
	code          []byte
	input         []byte
	iterations    []int
}

//NewInterpreter creates a new brainfuck interpreter using your specified cellcount and max iteration count
func NewInterpreter(cellcount, maxiter int) Interpreter {
	return Interpreter{make([]byte, cellcount), maxiter, 0, 0, nil, nil, nil}
}

func (I *Interpreter) goRight() {
	I.cellptr = (I.cellptr + 1) % len(I.cells)
}

func (I *Interpreter) goLeft() {
	I.cellptr--
	if I.cellptr < 0 {
		I.cellptr += len(I.cells)
	}
}

func (I *Interpreter) plus() {
	I.cells[I.cellptr]++
}

func (I *Interpreter) minus() {
	I.cells[I.cellptr]--
}

func (I *Interpreter) printChar() {
	fmt.Print(string(I.cells[I.cellptr]))
}

func (I *Interpreter) getChar() {
	if len(I.input) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		I.input = []byte(scanner.Text())
	}
	if len(I.input) != 0 {
		I.cells[I.cellptr] = I.input[0]
		I.input = I.input[1:]
	}
}

func (I *Interpreter) openLoop() {
	if I.cells[I.cellptr] != 0 {
		I.iterations = append(I.iterations, 0)
		return
	}
	openbr := 1
	for openbr > 0 {
		I.pc++
		switch I.code[I.pc] {
		case '[':
			openbr++
			break
		case ']':
			openbr--
			break
		}
	}
}

func (I *Interpreter) closeLoop() {
	if I.cells[I.cellptr] == 0 {
		I.iterations = I.iterations[:len(I.iterations) - 1]
		return
	}
	I.iterations[len(I.iterations) - 1]++
	if ! (I.iterations[len(I.iterations) - 1] < I.maxIterations) {
		fmt.Fprintln(os.Stderr, "FATAL ERROR: TOO MANY ITERATIONS")
		os.Exit(666)
	}
	closebr := 1
	for closebr > 0 {
		I.pc--
		switch I.code[I.pc] {
		case '[':
			closebr--
			break
		case ']':
			closebr++
			break
		}
	}
}

func (I *Interpreter) process() {
	switch I.code[I.pc] {
	case '>':
		I.goRight()
		break
	case '<':
		I.goLeft()
		break
	case '+':
		I.plus()
		break
	case '-':
		I.minus()
		break
	case '.':
		I.printChar()
		break
	case ',':
		I.getChar()
		break
	case '[':
		I.openLoop()
		break
	case ']':
		I.closeLoop()
		break
	}
}

//RunCode takes in brainfuckcode as a byte slice and executes it
func (I *Interpreter) RunCode(code []byte) {
	I.code = code
	for I.pc = 0; I.pc < len(I.code); I.pc++ {
		I.process()
	}
}
