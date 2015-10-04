package ui

import (
	"fmt"
	"regexp"

	"github.com/ukautz/clif"
)

var formatter *clif.DefaultFormatter
var output *clif.DefaultOutput
var input *clif.DefaultInput

func init() {
	formatter = clif.NewDefaultFormatter(clif.DefaultStyles)
	output = clif.NewOutput(nil, formatter)
	input = clif.NewDefaultInput(nil, output)
}

// Ask prompts the user to answer a question.
func Ask(question string, check func(string) error) string {
	return input.Ask(question, check)
}

// AskRegex promps the user to answer a question and checks the answer
// against the given regular expression.
func AskRegex(question string, rx *regexp.Regexp) string {
	return input.AskRegex(question, rx)
}

// Choose prompts the user to select one choice among many.
func Choose(question string, choices map[string]string) string {
	return input.Choose(question, choices)
}

// Printf prints a formatted message.
func Printf(msg string, args ...interface{}) {
	output.Printf(msg, args...)
}

// Errorf formats an error message and returns it.
func Errorf(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}
