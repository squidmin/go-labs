# Characteristics of Go

This section discusses the aspects that make Go a successful language in the programming world.

## Topics

- Features of Go
- Uses of Go
  - Used as a system programming language
  - Used as a general programming language
  - Used as an internal support
- Caveat
- Guiding design principles

## Features of Go

Go is essentially an **imperative** (procedural, structural) language, built with **concurrency** in mind.
It is not truly object-oriented like Java and C++ because it doesn't have the concepts of classes and inheritance.
However, it does have a concept of interfaces, with which much of the polymorphism can be implemented.
Go has a clear and expressive type-system, but it is lightweight and without hierarchy.
So in this respect, it could be called a **hybrid language**.

Some features of modern OOP languages were intentionally left out.
This is because object orientation was considered too _heavy_, often leading to cumbersome development constructing big type-hierarchies, and therefore non-compliant with the speed-related goals of the language.
As per the decision made by the Go team, the following OOP features are missing from Golang.
Although, some of them might still be implemented in its future versions.

- To simplify the design, no function or operator overloading was added.
- Implicit conversions were excluded to avoid the many bugs and confusion arising from this in languages like C / C++.
- No classes and type inheritance is supported in Golang.
- Golang does not support variant types. However, almost the same functionality is realized through interfaces.
- Dynamic code loading and dynamic libraries are excluded.
- Generics are not included (introduced in Go 1.18).
- Exceptions are not included (although recover / panic often does in that direction).
- Assertions are not included.
- Immutable (unable to change) variables are excluded.

A discussion around these features by the Go team itself can be found in the [Go FAQ](https://go.dev/doc/faq).

Go is a **functional language**, meaning that functions are the basic building blocks, and their use is very versatile.
There is more information on Golang functions in later sections.

Go is a statically and strongly typed language, enhancing safety and efficiency by compiling to native code and disallowing implicit type conversions.
Despite this, it incorporates some dynamic typing features, making it attractive to programmers from both statically and dynamically typed language backgrounds.

Last but not least, Go has support for cross-compilation; for example, developing and compiling on a Linux machine for an application that will execute on Windows.
It is one of the first programming languages that can use UTF-8 not only in strings but also in program code.
Go is truly an international language because Go source-files are also in UTF-8.

## Uses of Go

There are many uses of Go. Following are some main uses of this language:

- As a system programming language
- As a general programming language
- As an internal support

#### Used as a system programming language

Go was designed for server-heavy environments like web servers and storage systems, excelling in high-performance distributed systems due to its productivity and ease in handling massive concurrency and event processing.
This makes it ideal for game server and IoT development.

### Used as a general programming language

Go is versatile for various programming tasks like text processing, frontend development, and scripting but is not ideal for real-time applications due to its garbage collection and automatic memory management features.

### Used as an internal support

Go is particularly well-suited for developing distributed systems -- applications that run across multiple computers in a network to achieve a common goal.

Within Google, Go has been adopted for critical backend services that require high performance and reliability.
Specifically, parts of Google Maps, a complex service that processes vast amounts of data to provide real-time mapping, navigation, and location-based information, are implemented using Go.
This showcases Go's capability to handle heavy-duty, distributed applications by leveraging its features such as concurrency support, efficient execution, and simplicity in dealing with networked services.

## Caveat

If you come into Go and have a background in other contemporary (mostly class or inheritance-oriented languages like Java, C#, Objective C, Python, or Ruby), then you should use caution.
You can fall into the trap of trying to program in Go as you did in your previous language.
However, Go is built on a different model.
So if you decide to move code from language X to Golang, you will most likely produce non-idiomatic code that works poorly overall.

You have to start over by _thinking_ in Go.
If you take a higher point of view and start analyzing the problem from within the Go mindset, often a different approach suggests itself, which leads to an elegant and idiomatic Go solution.

## Guiding design principles

Go aims to simplify coding by using only 25 keywords, leading to less typing, less clutter, and reduced complexity.
Its clean and concise syntax, along with an LALR(1) grammar that doesn't require a symbol table for parsing, results in faster compilation times.

These aspects reduce the number of code lines necessary, even when compared with a language like Java.
Additionally, Go has a _minimalistic approach_: there tends to be only one way of getting things done, so reading other people's code is generally pretty easy, and we all know the code's readability is of the utmost importance in software engineering.

The design concepts of the language are _orthogonal_ because they don't stand in each other's way, and they don't add up complexity to one another.

Golang is completely defined by an explicit specification that can be found [here](https://go.dev/ref/spec); it is not defined by an implementation as is Ruby, for example.
An explicit language specification was required for implementing the _two different compilers_ **gc** and **gccgo**, and this in itself was a great help in clarifying the specification.
