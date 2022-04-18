package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	tm "github.com/buger/goterm"
	"github.com/fatih/color"
)

// holds the basic windows commands for linux users
// if any argument contains /? show help page

func RegisterCommands() {
	RegisteredCommands["cd"] = Command{
		Execute: Cd,
		Help: func(args ...string) {

		},
		Description: "Changes the current working directory.",
		Usage:       "COLOR [attr]",
	}
	RegisteredCommands["cls"] = Command{
		Name:    "cls",
		Execute: Cls,
		Help: func(args ...string) {

		},
		Description: "Clear the terminal screen.",
		Usage:       "CLS",
	}
	RegisteredCommands["color"] = Command{
		Name:    "color",
		Execute: Color,
		Help: func(args ...string) {

		},
		Description: "Changes the color of the terminal foreground or background.",
		Usage:       "COLOR [attr]",
	}
	RegisteredCommands["exit"] = Command{
		Execute: Exit,
		Help: func(args ...string) {

		},
		Description: "Exits the shell.",
		Usage:       "EXIT [/B] [Exitcode]",
	}
}

func Cls(Args ...string) {
	tm.Flush()
	tm.Clear()
	tm.MoveCursor(1, 1)
}

func Color(Args ...string) {
	if len(Args) == 1 {
		if len(Args[0]) >= 1 && len(Args[0]) <= 2 {
			if strings.ContainsAny(strings.ToUpper(Args[0]), "0123456789ABCDEF") {
				// changes the color of the terminal foreground and the color of the terminal background
				BackgroundTurn := false
				SetBackground := false
				fmt.Println([]byte(Args[0]))
				for _, v := range strings.ToUpper(Args[0]) {
					if BackgroundTurn {
						SetBackground = true
					}
					if v == '0' {
						if BackgroundTurn {
							color.Set(color.BgBlack)
						} else {
							color.Set(color.FgBlack)
						}
					} else if v == '1' {
						if BackgroundTurn {
							color.Set(color.BgBlue)
						} else {
							color.Set(color.FgBlue)
						}
					} else if v == '2' {
						if BackgroundTurn {
							color.Set(color.BgGreen)
						} else {
							color.Set(color.FgGreen)
						}
					} else if v == '3' {
						if BackgroundTurn {
							color.Set(color.BgCyan)
						} else {
							color.Set(color.FgCyan)
						}
					} else if v == '4' {
						if BackgroundTurn {
							color.Set(color.BgRed)
						} else {
							color.Set(color.FgRed)
						}
					} else if v == '5' {
						if BackgroundTurn {
							color.Set(color.BgMagenta)
						} else {
							color.Set(color.FgMagenta)
						}
					} else if v == '6' {
						if BackgroundTurn {
							color.Set(color.BgYellow)
						} else {
							color.Set(color.FgYellow)
						}
					} else if v == '7' {
						if BackgroundTurn {
							color.Set(color.BgWhite)
						} else {
							color.Set(color.FgWhite)
						}
					} else if v == '8' {
						if BackgroundTurn {
							color.Set(color.BgHiBlack)
						} else {
							color.Set(color.FgHiBlack)
						}
					} else if v == '9' {
						if BackgroundTurn {
							color.Set(color.BgHiBlue)
						} else {
							color.Set(color.FgHiBlue)
						}
					} else if v == 'A' {
						if BackgroundTurn {
							color.Set(color.BgHiGreen)
						} else {
							color.Set(color.FgHiGreen)
						}
					} else if v == 'B' {
						if BackgroundTurn {
							color.Set(color.BgHiCyan)
						} else {
							color.Set(color.FgHiCyan)
						}
					} else if v == 'C' {
						if BackgroundTurn {
							color.Set(color.BgHiRed)
						} else {
							color.Set(color.FgHiRed)
						}
					} else if v == 'D' {
						if BackgroundTurn {
							color.Set(color.BgHiMagenta)
						} else {
							color.Set(color.FgHiMagenta)
						}
					} else if v == 'E' {
						if BackgroundTurn {
							color.Set(color.BgHiYellow)
						} else {
							color.Set(color.FgHiYellow)
						}
					} else if v == 'F' {
						if BackgroundTurn {
							color.Set(color.BgHiWhite)
						} else {
							color.Set(color.FgHiWhite)
						}
					} else {
						fmt.Println("ALARM!")
						CommandHelp("color")
						return
					}
					BackgroundTurn = true
				}
				if !SetBackground {
					color.Set(color.BgBlack)
				}
			} else {
				CommandHelp("color")
			}
		}
	} else {
		// set default
		color.Set(color.FgWhite, color.BgBlack)
	}
}

func Cd(Args ...string) {
	if len(Args) == 0 {
		CurrentPath, Err := os.Getwd()
		if Err != nil {
			log.Fatalln(Err.Error())
		}
		fmt.Println(CurrentPath + "\n")
	} else if len(Args) == 1 {
		if Args[0] == ".." {
			CurrentPath, Err := os.Getwd()
			if Err != nil {
				log.Fatalln(Err.Error())
			}
			// get parent of current path
			Parent := filepath.Dir(CurrentPath)
			os.Chdir(Parent)
		} else {
			os.Chdir(Args[0])
		}
	} else {
		fmt.Println("The syntax for the file names, directory names or the drive designation is incorrect.")
	}
}

func Exit(Args ...string) {
	os.Exit(0)
}
