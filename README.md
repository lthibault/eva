# Eva

Exploratory implementation of the [Eva virtual machine](https://youtu.be/7pLCpN811tQ) in pure Go.

Eva is a simple virtual machine designed for educational use.  This is not intended to be a production-ready language implementation.

## Installation

`go get github.com/lthibault/eva`

Eva requires Go >= 17.0.0 with modules enabled.

## Features

The standard Eva VM, as designed by [Dmitry Soshnikov](https://github.com/dmitrysoshnikov), boasts the following features:

- Object-Oriented Programming
- Bytecode-emitting compiler
- 

Porting Eva to Go requires workarounds for some C++ features, such as tagged unions, leading to the development of additional features.  In addition, certain language features have been added for practice and exploration.

- Dynamic Values via [NaN-boxing](https://piotrduperas.com/posts/nan-boxing)
- Lightweight concurrency via Goroutines (planned)
- [Algebraic Effect Handlers](https://www.eff-lang.org/handlers-tutorial.pdf) (planned)
- [Lazy heap allocation](https://www.cs.tufts.edu/~nr/cs257/archive/henry-baker/cons-lazy-alloc.pdf) (planned)
