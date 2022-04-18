package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

type Command struct {
	Name        string               // "color"
	Help        func(Args ...string) // "full help page (customizable for arguments)"
	Execute     func(Args ...string) // "full execution function (customizable for arguments)"
	Description string               // "Changes the default background and foreground colors of the terminal."
	Usage       string               // "COLOR [attr]"
}

type JumpMark struct {
	Name     string    // ":main"
	Code     string    // "echo lol"
	NextMark *JumpMark /*

		:main
		 :: Jumps directly to the next mark
		:next

	*/
	NextCode string /*

		:main
		 :: Jumps to the next code
		echo lol
	*/
}

var PublicJumpMarks = make(map[string]JumpMark)

// Starting Code -> Jump Mark -> Next Code / Next Jump Maek

func CommandHelp(CommandName string) {
	// get command
	Command := RegisteredCommands[CommandName]
	// print help
	fmt.Println(Command.Help)
}

var RegisteredCommands = map[string]Command{}

func GetReleaseHash() string {
	// get file content of this current file
	FileContent, Err := os.Open(os.Args[0])
	if Err != nil {
		fmt.Println("Error:", Err)
		os.Exit(1)
	}
	// read the file content
	FileContentBytes, Err := ioutil.ReadAll(FileContent)
	if Err != nil {
		fmt.Println("Error:", Err)
		os.Exit(1)
	}
	// return the hash of the file content
	H := md5.New()
	H.Write(FileContentBytes)
	return fmt.Sprintf("%x", H.Sum(nil))
}

func main() {
	color.Set(color.FgWhite, color.BgBlack)
	// remove first letter of the OS
	OS := runtime.GOOS[1:]
	OS = strings.ToUpper(string(runtime.GOOS[0])) + OS
	fmt.Println("BatchPlus ["+ReleaseVersion+" "+ReleaseType+" - "+GetReleaseHash()+"] running on", OS, "("+runtime.GOARCH+")")
	fmt.Println("released under Creative Commons Zero v1.0 Universal made by x7t8c team")
	fmt.Println("")
	// println out the current path + ">"
	CurrentPath, Err := os.Getwd()
	if Err != nil {
		fmt.Println("Error:", Err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	RegisterCommands() // register all build-in commands

	for {
		fmt.Print(CurrentPath + ">")
		// wait for input from user
		Input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		Input = strings.Replace(Input, "\n", "", -1)
		NewInput := ""
		// iterate over the input and add it to the new input except 13
		for i := 0; i < len(Input); i++ {
			if Input[i] != 13 {
				NewInput += string(Input[i])
			}
		}
		// set input to new input
		Input = NewInput
		Command := strings.Split(Input, " ")[0]
		// remove command as prefix from input
		Args := strings.Split(Input, " ")
		// remove command from args
		Args = Args[1:]
		Interpret(Command, Args...)
	}
}
