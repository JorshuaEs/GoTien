# GOTIEN Compilador

Este proyecto implementa un intérprete para un lenguaje de programación personalizado similar, escrito en Go. El sistema incluye todas las etapas del proceso de interpretación: lexing, parsing, compilación y evaluación.

## Features

* Implementacion del lenguaje
    * variables con **let**
    * Expresiones aritméticas
    * Flujo de control (if/else)
    * Funciones y closures
    * Arrays y hash maps
    * funciones integradas
* Herramientas de compilado
    *   Análisis léxico
    *   Análisis sintáctico
    *   Compilador a lenguaje intermedio
    *   Máquina virtual basada en pila
    *   Interactive Read-Evaluate-Print Loop
* Features Avanzadas
    *   Closure support
    *   First-class functions
    *   Manejo de errores

## USO

### Modo REPL

```bash
$ gotien
>> let x = 5;
>> let square = fn(n) { n * n };
>> square(x);
25
```

### Compiado 

```bash 
$ ./gotien program.gt 
```
**Programa de ejemplo**
```javascript
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

puts(fibonacci(10));
```
**Resultado**
```bash
$55
``` 

## Syntax del lenguaje

### variables

```javascript
let x = 5;
let name = "GoTien";
let active = true;
```

### Funciones

```javascript
let add = fn(a, b) {
    return a + b;
};

let apply = fn(f, x, y) {
    return f(x, y);
};

apply(add, 2, 3); // 5
```
### Flujo de control

```javascript
let max = fn(a, b) {
    if (a > b) {
        return a;
    } else {
        return b;
    }
};
```

### Estructura de datos

```javascript
// Arrays
let nums = [1, 2, 3];
push(nums, 4);

// Hashes
let person = {
    "name": "Alice",
    "age": 30
};
person["name"]; // "Alice"
```
### Funciones integradas

| Función  | Descripción                        | Ejemplo                 |
|----------|------------------------------------|-------------------------|
| `len`    | Obtiene el tamaño del Array o String| `len([1,2,3])` → `3`    |
| `puts`   | Imprime valores                    | `puts("Hello")`         |
| `first`  | Obtiene el primer elemento del Array| `first([1,2,3])` → `1`  |
| `last`   | Obtiene el último elemento del Array| `last([1,2,3])` → `3`   |
| `rest`   | Obtiene el Array sin el primer elemento| `rest([1,2,3])` → `[2,3]` |
| `push`   | Agrega el elemento al Array                   | `push([1,2], 3)` → `[1,2,3]` |

## Descripción general de la arquitectura
```
Source Code → Lexer → Tokens → Parser → AST (Abstract Syntax Tree) → Compiler → Bytecode → VM → Execution
```
### Componentes clave
1. Lexer: Convierte el código fuente en tokens.
2. Parser: Construye el AST usando Pratt parsing.
3. Compilador: Genera el codigo intermedio utilizando el AST.
4. VM: Ejecuta el codigo intermedio con:
    * Operaciones de pila
    * Manejo de Estructura
    * Soporte de Closure

# Componentes del proyecto

1. Análisis Léxico `(lexer.go)`
Convertierte el código fuente en tokens.

*   Identifica números, strings, identificadores y símbolos especiales.
*   Maneja operadores de 1 y 2 caracteres (como ==, !=).
*   Ignora los espacios en blanco
*   Soporta strings entre comillas

2.  Árbol de Sintaxis Abstracta `(ast.go)`
Define la estructura en forma de árbol, reflejando la estructura del programa por medio de nodos, cada nodo representando una costrucción del lenguaje.

*   Estructura
    *   Nodo: Base común para todos los nodos.
        *   `TokenLiteral()`: Obtiene el token asociado
        *   `String()`: Representación en cadena del nodo

**Ejemplo**

Código fuente:
```javascript
let x = 5;
```
se convierte en:

```go
&LetStatement{
	Token: token.LET,
	Name:  &Identifier{Value: "x"},
	Value: &IntegerLiteral{Value: 5},
}
```
3. Sistema de Instrucciones Bytecode `(code.go)`
Formato intermedio de Instrucciones que representa las operaciones de un lenguaje de porgramacion. Las instrucciones se cofican como lenguaje intermedio, este paquete ayuda a definir, crear, leer y formatear las instrucciones. 

Este conmponente contiene:
*   Definiciones de los opcode (códigos de operación).
*   Funciones para crear instrucciones (`make`).
*   Funciones para leer operandos de instrucciones(`ReadOperands`).
*   Funcionalidad para convertir las instrucciones a un formato humano legible (`Instructions.String()`).

**Ejemplo**

Se traduce una expresión como 1 + 2 a bytecode (`OpConstant`, `opConstant`, `OpAdd`)
* Un ejemplo más específico
Se tiene la instruccion para cargar una constante a la memoria.

```go
ins := Make(OpConstant, 65534)
fmt.Println(Instructions(ins).String())

```
Se traduce la instrucción a código intermedio como:

```yaml
0000 OpConstant 65534
```
4. Compilador `(compiler.go)`
Transforma el AST en código intermedio, el cual puede ser ejecutado por la vm. Recorre y visita todos los nodos del árbol, emite instrucciones en código intermedio según el tipo de nodo (como literales, operadores, condicionales, funciones, etc.).

**Pasos del compilador**

1. Analiza el AST.
2. Convierte nodos en bytecode.
3. Guarda constantes como enteros o funciones en una lista.
4. Usa una tabla de símbolos para gestionar variables y ámbitos.
5. Emite instrucciones como:
    * OpConstant (cargar una cosntante).
    * OpAdd, OpSub, etc. (operaciones aritméticas).
    * OpPop (descartar valor del stack).
    * OpJum, OpJumNotTruthy (saltos condicionales).

**Ejemplo**
```go
let x = 1 + 2;
```
El compilador genera algo como:
```assembly
OpConstant 1
opConstant 2
OpAdd
OpSetGlobal 0 //x
```