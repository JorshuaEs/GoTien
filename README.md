[⚠️ Suspicious Content] 
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
