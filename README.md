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
| `len`    | Get length of array or string      | `len([1,2,3])` → `3`    |
| `puts`   | Print values                       | `puts("Hello")`         |
| `first`  | Get first array element            | `first([1,2,3])` → `1`  |
| `last`   | Get last array element             | `last([1,2,3])` → `3`   |
| `rest`   | Get array without first element    | `rest([1,2,3])` → `[2,3]` |
| `push`   | Append to array                    | `push([1,2], 3)` → `[1,2,3]` |

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