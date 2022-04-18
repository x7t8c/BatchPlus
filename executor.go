package main

import (
	"fmt"
	"os"
)

func main() {
	// if not arguments are passed
	fmt.Println(ParseEnvVarsAndConvert("%PAPA% ist der Pfad und %MAMA% auch!"))
	if len(os.Args) == 1 {
		fmt.Println("No arguments passed, displaying help")
		os.Exit(1)
	} else {
		Interpret(os.Args[1], os.Args[2:]...)
	}
}
