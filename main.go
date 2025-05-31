/*
package main

import (

	"fmt"
	"gotien/repl"
	"os"
	"os/user"

)

	func main() {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the GoTien programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
*/
/*package main

import (
	"fmt"
	"gotien/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gotien/compiler"
	"gotien/lexer"
	"gotien/parser"
	"gotien/vm"
)

// The main function reads a file, parses its content, compiles it into bytecode, runs it on a virtual
// machine, and prints the result.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: gotien archivo.gotien")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error leyendo archivo: %s\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Errores encontrados durante el análisis:")
		for _, err := range p.Errors() {
			fmt.Printf(" - %s\n", err)
		}
		os.Exit(1)
	}

	comp := compiler.New()
	err = comp.Compile(program)
	if err != nil {
		fmt.Printf("Error compilando: %s\n", err)
		os.Exit(1)
	}

	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		fmt.Printf("Error en VM: %s\n", err)
		os.Exit(1)
	}

	result := machine.LastPoppedStackElem()
	fmt.Println(result.Inspect())
}
