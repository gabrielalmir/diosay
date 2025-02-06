package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var message string

	if len(os.Args) > 1 {
		message = strings.Join(os.Args[1:], " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Erro ao ler entrada:", err)
			os.Exit(1)
		}
	}

	if message == "" {
		fmt.Fprintln(os.Stderr, "Uso: echo \"mensagem\" | diosay ou diosay <mensagem>")
		os.Exit(1)
	}

	messageBox := createMessageBox(message)
	logo := `
     -=+++++++-
   =+++++*******+=
 -++++=======+****=
=++++------.. .:***=
=+++***..=***:  =**+
+++***..+****- .+***
=+***-.:***==..+***+
=+**=..+:=-..-*****+
 -+*=====++*******+
  -+*************+
     =+*******+
`

	fmt.Println("\n" + messageBox)
	fmt.Println("    \\")
	fmt.Println("     \\")
	fmt.Println(logo)
}

func createMessageBox(message string) string {
	border := strings.Repeat("-", len(message)+4)
	top := fmt.Sprintf(" %s", border)
	middle := fmt.Sprintf("|  %s  |", message)
	bottom := fmt.Sprintf(" %s", border)

	return fmt.Sprintf("%s\n%s\n%s", top, middle, bottom)
}
