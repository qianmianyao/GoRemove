package main

import (
	"grm/mainFunc"
	"os"
)

func main() {
	file := os.Args
	mainFunc.Command(&file)
}
