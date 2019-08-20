package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/ivpusic/monkey/evaluator"
	"github.com/ivpusic/monkey/lexer"
	"github.com/ivpusic/monkey/object"
	"github.com/ivpusic/monkey/parser"
	"github.com/ivpusic/monkey/repl"
)

func startRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func interpretFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("error while parsing monkey source file. %s", err.Error()))
	}

	env := object.NewEnvironment()

	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
}

func main() {
	if len(os.Args) > 1 {
		interpretFile(os.Args[1])
	} else {
		startRepl()
	}
}
