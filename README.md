# hello-world-go

hello world made with golang following the documentation

## But is not just that

### Documentation

- [x] Create a module -> Write a small module with functions you can call from another module.
- [x] Call your code from another module -> Import and use your new module.
- [x] Return and handle an error -> Add simple error handling.
- [x] Return a random greeting -> Handle data in slices (Go's dynamically-sized arrays).
- [x] Return greetings for multiple people -> Store key/value pairs in a map.
- [x] Add a test -> Use Go's built-in unit testing features to test your code.
- [x] Compile and install the application -> Compile and install your code locally.

### A Tour of Go

#### Part 1

- [x] [Hello](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#hello)
- [x] [Sandbox](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#sandbox)
- [x] [Packages](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#packages)
- [x] [Imports](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#imports)
- [x] [Exported-names](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#exported-names)
- [x] [Functions](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#functions)
- [x] [Functions Continued](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#functions-continued)
- [x] [Multiple Results](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#multiple-results)
- [x] [Named Results](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#named-results)
- [x] [Variables](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#variables)
- [x] [Variables With Initializers](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#variables-with-initializers)
- [x] [Short Variable Declarations](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#short-variable-declarations)
- [x] [Basic Types](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#basic-types)
- [x] [Zero](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#zero)
- [x] [Type Conversions](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#type-conversions)
- [x] [Type Inference](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#type-inference)
- [x] [Constants](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#constants)
- [x] [Numeric Constants](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#numeric-constants)

#### Part 2

- [x] [For](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#for)

### Hello

```
These example programs demonstrate different aspects of Go. The programs in the tour are meant to be starting points for your own experimentation.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Sandbox

```
This tour is built atop the Go Playground, a web service that runs on golang.org's servers.

The service receives a Go program, compiles, links, and runs the program inside a sandbox, then returns the output.

There are limitations to the programs that can be run in the playground:

In the playground the time begins at 2009-11-10 23:00:00 UTC (determining the significance of this date is an exercise for the reader). This makes it easier to cache programs by giving them deterministic output.
There are also limits on execution time and on CPU and memory usage, and the program cannot access external network hosts.
The playground uses the latest stable release of Go.
```

Read "[Inside the Go Playground](https://go.dev/blog/playground)" to learn more.

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Packages

```
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

(To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Imports

```
This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:
```

```go
import "fmt"
import "math"
```

```
But it is good style to use the factored import statement.

The right way is:
```

```go
import (
	"fmt"
	"math"
)
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Exported Names

```
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Functions

```
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.
```

For more about why types look the way they do, see the [article on Go's declaration syntax](https://blog.golang.org/gos-declaration-syntax).

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Functions Continued

```
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened
```

```go
x int, y int
```

to

```go
x, y int
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Multiple Results

```
A function can return any number of results.

The swap function returns two strings.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Named Results

```
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Variables

```
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Variables With Initializers

```
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Short Variable Declarations

```
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Basic Types

```
Go's basic types are
```

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

```
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Zero

```
Variables declared without an explicit initial value are given their zero value.

The zero value is:

  0 for numeric types,
  false for the boolean type, and
  "" (the empty string) for strings.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Type Conversions

```
The expression T(v) converts the value v to the type T.

Some numeric conversions:
```

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

```
Or, put more simply:
```

```go
i := 42
f := float64(i)
u := uint(f)
```

```
Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Type Inference

```
When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

When the right hand side of the declaration is typed, the new variable is of that same type:
```

```go
var i int
j := i // j is an int
```

```
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:
```

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

```
Try changing the initial value of v in the example code and observe how its type is affected.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Constants

```
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### Numeric Constants

```
Numeric constants are high-precision values.

An untyped constant takes the type needed by its context.

Try printing needInt(Big) too.

(An int can store at maximum a 64-bit integer, and sometimes less.)
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)

### For

```
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

  the init statement: executed before the first iteration
  the condition expression: evaluated before every iteration
  the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
```

[Go back to "A Tour of Go" menu](https://github.com/DPechetti/hello-world-go/tree/feature/tour-of-go#a-tour-of-go)
