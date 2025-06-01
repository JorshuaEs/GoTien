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

1. Análisis Léxico `(lexer.go)`\
Convertierte el código fuente en tokens.

*   Identifica números, strings, identificadores y símbolos especiales.
*   Maneja operadores de 1 y 2 caracteres (como ==, !=).
*   Ignora los espacios en blanco
*   Soporta strings entre comillas

**Ejemplo**

Se tiene el siguiente código:

```javascript
let x = 5;
```

El lexer devuelve lo siguiente: 

```go
{Type: "LET", Literal: "let"}
{Type: "IDENT", Literal: "x"}
{Type: "ASSIGN", Literal: "="}
{Type: "INT", Literal: "5"}
{Type: "SEMICOLON", Literal: ";"}
```

2.  Árbol de Sintaxis Abstracta `(ast.go)`\
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
3. Sistema de Instrucciones Bytecode `(code.go)`\
Formato intermedio de Instrucciones que representa las operaciones de un lenguaje de porgramacion. Las instrucciones se codifican como bytecode, este paquete ayuda a definir, crear, leer y formatear las instrucciones. 

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
Se traduce la instrucción a bytecode como:

```yaml
0000 OpConstant 65534
```
4. Compilador `(compiler.go)`\
Transforma el AST en bytecode, el cual puede ser ejecutado por la vm. Recorre y visita todos los nodos del árbol, emite instrucciones en bytecode según el tipo de nodo (como literales, operadores, condicionales, funciones, etc.).

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
5. Tabla de símbolos\
Encargado de gestionar y resolver nombres de variables y funcones segun su alcance durante la compilación.
Implementa una tabla de símbolos con somporte para múltiples scopes, trabajando estrechamente con el compilador para rastrear qué significa cada nombre que aparece en el código fuente, y en que contexto se debe interpretar.

**Ejemplo**

Se tiene el siguiente código  para compilar: 
```javascript
let a = 5;

fn outer() {
  let b = 10;

  fn inner() {
    return a + b;
  }

  return inner();
}
```
Resultado final de la tabla de simbolo

```sql
Global scope:
  - a: GlobalScope, Index 0
  - outer: GlobalScope, Index 1

outer scope:
  - b: LocalScope, Index 0
  - inner: LocalScope, Index 1

inner scope:
  - b: FreeScope, Index 0
  - a: GlobalScope (resuelto desde global)
```
6.  Funciones Integradas (builtins.go)/
Definición de un conjunto de funciones integradas, similares a las funciones estándar qye ofrecen lenguajes como Python o JavaScript.

Funciones disponibles automáticamente desntro del lenguaje, sin necesida de que el usuario las defina. Funciones integradas definidas: 
* len → Devuelve la longitud de una cadena o array.
* puts → Imprime en consola los argumentos.
* first → Devuelve el primer elemento de un array.
* last → Devuele el último elemento de un array.
* rest → Devuelce un nuevo array con todos los elementos excepto el primero.
* push → Devuelve un nuevo array con el segundo argumento añadido  al final del primero.

7. Entorno `(Environment.go)`\

Es donde se guardan y consultan variables durante la ejecución  del programa. Es un tipo de estructura de datos que actúa como un diccionario, donde se almacenan pares nombre → valor.

**Ejemplo**

```go
x = 5
y = "hola"
```
8. Tipos de Objectos `(object.go)`\

Define todos los tiepos de objetos que pueden existir en el lenguaje. Representa los valores del lenguaje tal como los vería el compilador en tiempo de ejecución. Entre los mas importantes se encuentran; `Interger`,`Boolean`, `String`.

9. Analizador léxico `(parser.go)`\

Tiene como propósito  transformar una secuencia de tokesn (producida por el lexer) en una estructura de árbol (AST) que representa el programa en sí. Toma como entrada una lista de tokens que viene del lexer y construye un AST, que modela el programa como una estructura jerárquica  que el compilador puede ejecutar.

**Ejemplo**

se tiene lo siguiente: `let x = 5 + 2`;

* Tokens obtenidos del lexer:

```javascript
LET IDENT ASSIGN INT PLUS INT SEMICOLON
```

* El AST que genera el parser:

```javascript 
&ast.Program{
  Statements: []ast.Statement{
    &ast.LetStatement{
      Name: "x",
      Value: &ast.InfixExpression{
        Left:  &ast.IntegerLiteral{Value: 5},
        Op:    "+",
        Right: &ast.IntegerLiteral{Value: 2},
      },
    },
  },
}
```

10. Tokens `(tokens.go)`\

Define el sistema de tokens utilizados. Los tokes son las unidades léxicas que el lexer detecta al leer el código fuente. Es una tabla de referencia que otros compoentes usan para entender el código fuente. Se agrupan en las siguientes categorias:
* Tokens especiales
```go 
ILLEGAL = "ILLEGAL" // carácter no reconocido
EOF     = "EOF"     // fin de archivo
```
* Identificadores y literales
```go
IDENT, INT, STRING
```
* operadores
```go 
=, +, -, !, *, /, <, >, ==, !=
```
* Delimitadores
```go
( ) { } [ ] , ; :
```
* Palabras clave del lenguaje
```go
FUNCTION, LET, TRUE, FALSE, IF, ELSE, RETURN
```

10. Máquina Virtual (VM) `(vm.go)`\
implementación de una VM para la ejecución del bytecode generado por el compilador. Simula la ejecución del programa previamente compilado, que contiene instrucciones.
*   Caracteristicas
    * Es una VM basada en pila.
    * Se apoya en estructuras del compilador (`compiler`), instrucciones (`code`) y objetos del lenguaje (`object`).

**Ejemplo**

* Código fuente

```javascript
let result = 1 + 2;
```
* Bytecode
```assembly
OpConstant 0   // Push 1
OpConstant 1   // Push 2
OpAdd          // Add top two values (1 + 2)
OpPop          // Pop result (3)
```
* Ejecución de la VM
* Paso 1

```go
stack = [1]
sp = 1
```

* Paso 2

```go
stack = [1, 2]
sp = 2
```

* Paso 3

```go
stack = [3]
sp = 1
```

* Paso 4

```go
stack = []
sp = 0
EvaluatedValues = [3]
```

* Resultado

```go
3
```

11. Principal `(main.go)`\

Se encarga de compilar y ejecutar el código escrito por el usuario. Este programa:

* Lee y analiza el código fuente.
* Lo convierte a bytecode.
* Lo ejecuta en una VM.
* Imprime el resultado.