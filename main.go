package main

import (
	"fmt"
	"os"

	"gotien/compiler"
	"gotien/lexer"
	"gotien/parser"
	"gotien/repl"
	"gotien/vm"
)

func main() {
	if len(os.Args) < 2 {
		// Si no hay archivo: iniciar el REPL
		fmt.Println("Bienvenido a GoTien - modo interactivo (REPL)")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Si se proporciona un archivo: ejecutar archivo
	filename := os.Args[1]
	content, err := os.ReadFile(filename)
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
