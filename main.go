// Description: cowsay but using Golang
// References:
//	https://flaviocopes.com/go-tutorial-cowsay

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error!")
		return
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortne | gophersay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
}
