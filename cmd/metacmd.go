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

package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var usageText string

type MetaCmd struct {
	Name string
	Args []string
}

func init() {
	file, _ := os.Open("cmd/usage.txt")
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, _ := file.Read(buffer)

		if n == 0 {
			break // reached EOF
		}
	}

	usageText = string(buffer)
}

func NewMetaCmd(input string) *MetaCmd {
	arr := strings.Split(input[1:], " ")

	name := strings.ToLower(arr[0])
	var args []string
	if len(arr) > 1 {
		args = arr[1:]
	}

	return &MetaCmd{
		Name: strings.ToLower(name),
		Args: args,
	}
}

func (m *MetaCmd) Execute() error {
	switch m.Name {
	case "help":
		printUsage()
		return nil
	case "load":
		println("load is executed with filename ", m.Args[0])
		return nil
	case "flush":
		println("flush is executed with filename ", m.Args[0])
		return nil
	case "quit":
		println("quit is executed with args ")
		return nil
	default:
		return errors.New("invalid command. Type .help or /help for all meta commands")
	}
}

func (m *MetaCmd) ToString() string {
	return fmt.Sprintf("[MetaCmd] { Name: %s, Args : %v }", m.Name, m.Args)
}

func printUsage() {
	fmt.Println(usageText)
}
