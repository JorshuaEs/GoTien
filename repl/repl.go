package repl

import (
	"bufio"
	"fmt"
	"gotien/compiler"
	"gotien/lexer"
	"gotien/object"
	"gotien/parser"
	"gotien/vm"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)

	symbolTable := compiler.NewSymbolTable()
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")
	}
}

const GOTIEN = `       		 __,__
		GoTien
               ╭──────────────╮
               │   ▄───▄      │
         ╭─────┴──────────────┴────╮
         │     ╭────────────╮      │
         │     │   ▄────▄   │      │
   ╭─────┴─────┴────────────┴──────┴─────╮
   │         ╭────────────────╮          │
   │         │   ▄──────▄     │          │
   │         │   █      █     │          │
   │         │   █      █     │          │
   │         │   █▄▄▄▄▄▄█     │          │
   ╰─────────────────────────────────────╯
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, GOTIEN)
	io.WriteString(out, "Oh no... an ancient mist has clouded your code.\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
