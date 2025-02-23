package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// list of quad executables

var quads = []string{"./quadA", "./quadB", "./quadc", "./quadD", "./quadE"}

// run executes the quad checker

func Run() {
	// read input shape
	input, err := io.ReadAll(os.Stdin) // read all inputs() from the command line and store them in the input
	if err != nil || len(input) == 0 {
		fmt.Println("invalid input")
		return
	}
	// remove the extra new lines
	shape := input
	if input[len(input)-1] == '\n' {
		shape = input[:len(input)-1] // shape now ends at the last index excluding the new line
	}
	var width, height int
	for i, b := range input {
		if b == '\n' {
			if width == 0 {
				width = i
			}
			height++
		}
	}

	if width == 0 {
		width = len(input)
		height = 1
	}

	// store matches

	var matches []byte

	//compare input with each quad program
	for _, quad := range quads {
		output, err := exec.Command(quad, fmt.Sprintf("%d", width), fmt.Sprintf("%d", height)).Output()
		if err != nil {
			continue // skip if execution fails
		}
		// Trim newline from output
		if len(output) > 0 && output[len(output)-1] == '\n' {
			output = output[:len(output)-1]
		}
		// compare input with quad output
		if len(output) == len(shape) {
			same := true
			for i := range output {
				if output[i] != shape[i] {
					same = false
					break
				}
			}
			if same {
				if len(matches) > 0 {
					matches = append(matches, " || "...)
				}
				matches = append(matches, fmt.Sprintf("%s [%d] [%d]", quad[2:], width, height)...) // remove "./"
			}
		}
	}
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(string(matches))
	}
}

func main() {
	Run()
}
