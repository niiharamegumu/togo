package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	stdinScanner *bufio.Scanner
)

func getStdinScanner() *bufio.Scanner {
	if stdinScanner == nil {
		stdinScanner = bufio.NewScanner(os.Stdin)
	}
	return stdinScanner
}

// InputMultiLine reads multiple lines of input from the user until two consecutive newlines are entered.
func InputMultiLine(prompt string) string {
	scanner := getStdinScanner()
	fmt.Println(prompt)
	fmt.Println("Press Enter twice to finish, you can enter multiple lines")

	var builder strings.Builder
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			break
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}

	return strings.TrimSpace(builder.String())
}
