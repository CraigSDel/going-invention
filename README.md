# GOing-invention

## Installing GoLang

* Download the installer package
    * [GO](https://golang.org/)

## Why was the language created?

The language was created in 2007 to improve programming productivity with the rise of multicore, networked machines and
larger code bases. The goals of the programming language

* Static typing and run-time efficiency
* Readability and usability
* High-performance networking and multiprocessing.

## What environment variables are created?

### GOPATH

* Variable that defines the root of your workspace

### GOROOT

* Variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use
  different GO versions.

## Data types

* Numeric
    * int
    * float
* Boolean
* String
* Derived
    * Pointer (Memory location of a variable)
    * Array (Collection of similar data types)
    * Structure (Creating custom data types)
    * Map
    * Interface (Used to create abstractions)

## Operators

* Arithmetic
    * 1 + 2 = 3 (addition)
    * 2 - 1 = 1 (subtraction)
    * 2 * 2 = 4 (multiplication)
    * 2 / 2 = 1 (division)
    * 5 % 2 = 1 (modulus)

* Relational
    * 3 > 1 (greater than)
    * 1 < 3 (less than)
    * 3 >= 2 (greater than or equal)
    * 3 == 3 (equivalence)
    * 3 != 2 (not equals)

* Logical
    * && (and)
    * || (or)
    * ! (negation)

## Control Structures

* Defer
    * Defers the execution of the statement until the surrounding function has completed
    * Multiple defers are pushed on a stack and is executed using in the last in first out sequence
    * Used to clean up resources like database connections, file handling etc

* Regain
    * Assists with getting control of normal execution after a panic
    * Used with defer statement to recover in a go routine

* Panic
    * Similar to throwing an exception
    * When a panic is called normal execution flow is interrupted
    * Deferred functions are executed normally 