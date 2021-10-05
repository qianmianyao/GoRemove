package main

import (
	"os"
	"remove_Go/mainFunc"
)

func main() {
	file := os.Args
	mainFunc.Command(&file)
}
