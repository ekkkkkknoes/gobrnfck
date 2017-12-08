package main

import (
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
		inter := NewInterpreter(32767,32767)
		inter.RunCode(bfCode)
		return
	}
	fmt.Println("Error: Must specify brainf*ck file!")
	return
}
