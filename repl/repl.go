// Copyright (c) 2023 Priyak Dey
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the “Software”),
// to deal in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies
// or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR
// PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/priyakdey/yadb/cmd"
	"github.com/priyakdey/yadb/version"
)

func Run() {
	printWelcomeMessage()

	for {
		fmt.Print("yadb > ")

		// create a reader for taking input from stdin
		r := bufio.NewReader(os.Stdin)
		var buf []byte

		buf, _ = r.ReadBytes('\n')

		trimLeadingSpaces(&buf)
		trimNewLineAtEnd(&buf)

		var c cmd.Cmd

		// if leading character is dot or slash, we accept it as meta command
		if buf[0] == byte(DOT) || buf[0] == byte(SLASH) {
			c = cmd.NewMetaCmd(string(buf))
		} else {
			// this is sql command, if it does not end with a a SEMICOLON, continue taking input
			takeSqlInput(r, &buf)
			fmt.Println("[Sql Command]", string(buf))
		}

		if err := c.Execute(); err != nil {
			fmt.Printf(Red+"error: %s\n"+Reset, err)
		}
	}
}

// helper function which helps takes sql commands from multiple lines
func takeSqlInput(r *bufio.Reader, buf *[]byte) {
	fmt.Printf("..... ")
	data, _ := r.ReadBytes('\n')

	trimLeadingSpaces(&data)
	trimNewLineAtEnd(&data)

	*buf = append(*buf, byte(SPACE))
	*buf = append(*buf, data...)

	end := len(data)
	if data[end-1] == byte(SEMICOLON) {
		return
	} else {
		takeSqlInput(r, buf)
	}
}

func trimLeadingSpaces(arr *[]byte) {
	count := 0
	end := len(*arr)

	for i := 0; i < end; i++ {
		if (*arr)[i] != 32 {
			break
		}
		count++
	}

	*arr = (*arr)[count:]
}

func trimNewLineAtEnd(arr *[]byte) {
	count := 0
	end := len(*arr)

	for i := end - 1; i >= 0; i-- {
		if (*arr)[i] != byte(LF) && (*arr)[i] != byte(CR) {
			break
		}

		count++
	}

	*arr = (*arr)[:end-count]
}

func printWelcomeMessage() {
	fmt.Printf("Welcome to yadb-%s.\n", version.Version)
	fmt.Println(Red + "WARNING: yadb is an education project and not meant to be used with projects" + Reset)
	fmt.Println("Type .help or /help for more details")
	fmt.Println()
}
