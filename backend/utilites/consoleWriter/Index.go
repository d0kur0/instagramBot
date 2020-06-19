package consoleWriter

import "gopkg.in/gookit/color.v1"

func Header(message string) {
	color.Cyan.Printf("\n⇝ %s\n", message)
}

func Info(message string) {
	color.New(color.FgWhite, color.BgBlack).Printf(" ⮡ %s\n", message)
}

func Success (message string) {
	color.Green.Printf(" ⮡ %s\n", message)
}

func Error (message string) {
	color.Red.Printf(" ⮡ %s\n", message)
}