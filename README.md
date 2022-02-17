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
- [x] multiple-results.go
- [x] named-results.go
- [x] variables.go
- [x] variables-with-initializers.go
- [x] short-variable-declarations.go
- [x] basic-types.go
- [x] zero.go
- [x] type-conversions.go
- [x] type-inference.go
- [x] constants.go
- [x] numeric-constants.go

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
