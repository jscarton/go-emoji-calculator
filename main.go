// example.go
package main

import (
	"bufio"
	"fmt"
	"go-emoji-calculator/calculator"
	"go-emoji-calculator/translator"
	"io/ioutil"
	"os"
)

// Helper providing the equivalent of fmt.Printf to STDOUT
func fmtErr(msg string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg, a...)
}

func main() {

	// Read in STDIN
	expression, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		fmtErr("Error: Unable to read from STDIN: %s\n", err.Error())
		return
	}

	if len(expression) > 0 {
		translated_expr := translator.EmojiToAscii(string(expression))
		result := translator.NumberToEmoji(calculator.Calc(translated_expr))
		fmt.Printf("%s\n", result)
	} else {
		fmt.Printf("Error: There is not an expression to parse\n\U0001f937\n")
	}

}
