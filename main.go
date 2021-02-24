// Description: cowsay but using Golang
// References:
//	https://flaviocopes.com/go-tutorial-cowsay

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// this function takes a slice of strings of max width
// prepends/appends margins on first & last line
// 	and at start/end of each line
// retuns string with contents of the balloon
func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var result []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	result = append(result, top)

	// creates borders around the text depending on certain rules
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		result = append(result, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		result = append(result, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			result = append(result, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		result = append(result, s)
	}

	result = append(result, bottom)
	return strings.Join(result, "\n")
}

// this function convets all tabs to 4 spaces
// this prevents formatting issues
func tabsToSpaces(lines []string) []string {
	var result []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		result = append(result, l)
	}
	return result
}

// calculate max width given several strings
func calculateMaxWidth(lines []string) int {
	width := 0
	for _, line := range lines {
		// gets the length of the string
		len := utf8.RuneCountInString(line)
		// sets max width
		if len > width {
			width = len
		}
	}

	return width
}

// makes all of the strings have the same number of spaces
func normalizeStringsLength(lines []string, maxwidth int) []string {
	var result []string
	for _, line := range lines {
		// appends spaces to the end of the string so all are length maxwidth
		s := line + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(line))
		result = append(result, s)
	}
	return result
}

// given a figure name, print the figure
func printFigure(name string) {

	// raw strings with ascii figures
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	var cat = `   \
      /-/\
     ('-' )
      |   \ /
      U-U(_/
  `

	var gopher = `      \
	  (\____/)
	  /      \
	(o ) __ ( o)
 	 |   UU   | 
	<|        |>
	 |        |
	  \______/
	   U    U
`

	var random = `      \
	  ( )___( )
	( O  -  O |
	 |        |
`

	switch name {
	case "gopher":
		fmt.Println(gopher)
	case "cow":
		fmt.Println(cow)
	case "cat":
		fmt.Println(cat)
	default:
		fmt.Println(random)
	}
}

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error!")
		return
	}

	// checks to make sure input was piped in
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gophersay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	// reads each line
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	lines = tabsToSpaces(lines)
	maxWidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxWidth)
	balloon := buildBalloon(messages, maxWidth)

	fmt.Println(balloon)

	var figure string
	flag.StringVar(&figure, "f", "gopher", "the figure name. Valid alternate values are `cow`, `cat`, and `random`")
	flag.Parse()

	printFigure(figure)
	fmt.Println()
}
