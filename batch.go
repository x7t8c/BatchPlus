package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseEnvVarsAndConvert(Code string) string {
	// converts %PATH% to C:\Users\User\AppData\Local\Temp
	// iterate over all symbole in code
	// if symbol is %
	// get the next symbol
	ReadingVar := false
	ReadVar := ""
	NewCode := ""
	for i := 0; i < len(Code); i++ {
		if Code[i] == '%' {
			if ReadingVar {
				Value, Exists := os.LookupEnv(ReadVar)
				if Exists {
					NewCode += Value
				} else {
					NewCode += "%" + ReadVar + "%"
				}
				ReadVar = ""
			}
			ReadingVar = !ReadingVar
		} else if ReadingVar {
			ReadVar += string(Code[i])
		} else {
			NewCode += string(Code[i])
		}
	}
	return NewCode
}

func Interpret(Command string, Args ...string) {
	// parse env vars and convert the command and all arguments
	Command = ParseEnvVarsAndConvert(Command)
	for i := 0; i < len(Args); i++ {
		Args[i] = ParseEnvVarsAndConvert(Args[i])
	}
	// execute the command
	LoweredCommand := strings.ToLower(Command)
	ExecutedAnything := false
	if LoweredCommand == "exit" {
		ExecutedAnything = true
		os.Exit(0)
	} else if LoweredCommand == "cd" {
		ExecutedAnything = true
		Cd(Args...)
	}
	if ExecutedAnything {
		ExecutedAnything = !ExecutedAnything
		fmt.Println("\n")
	}
}
