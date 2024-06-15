# Reasons for Developing Go

This section covers some reasons why there was a need for the Go language's development.

## Topics

- Languages that influenced Go
- Why a new language?
    - Evolving with the computing landscape
    - Need for faster software development
    - Need for efficiency and ease of programming
- Targets of Go
    - Support for network communication, concurrency, and parallelization
    - Support for excellent building speed
    - Support for memory management

## Languages that influenced Go

Go, also known as Golang, belongs to the C-family, like C++, Java, and C#, and is inspired by many other languages created and used by its designers.
The resemblance with the syntax of the C language was maintained to gain an advantage of familiarity among developers.
The syntax of C was refined to be less verbose and complex.
However, compared to C / C++, the syntax is made more incisive (concise).
Additionally, Go has the features of a dynamic language, so Python and Ruby programmers feel more comfortable with it in terms of ease of use, simplicity, and efficient handling of high-level tasks.

## Why a new language?

Following are the reasons that lef to the development of Go:

- Evolving with the computing landscape
- Need for faster software development
- Need for efficiency and ease of programming

### Evolving with the computing landscape

Programming languages like C / C++ did not evolve with the computing landscape, so there is a need for a new systems language, appropriate for the needs of our computing era.

### Need for faster software development

Despite advancements in computing power, software development hasn't seen proportional improvements in speed or success rates, with many projects failing even as applications become larger.
This situation highlighted the need for a new low-level language that incorporates advanced concepts to address these challenges.

### Need for efficiency and ease of programming

Before Go, a developer had to choose between fast execution but slow and inefficient building (like C++) or efficient compilation but not so fast execution (like .NET or Java), or ease of programming but slower execution (e.g., dynamic languages such as Python, Ruby, or JavaScript).
Go is an attempt to combine these useful traits of languages: efficient and fast compilation, fast execution, and ease of programming.

## Targets of Go

The main target of Golang's design was to combine the efficacy, speed, and safety of a statically typed and compiled language with the ease of programming of a dynamic language to make programming more enjoyable.

Some other targets that Go was meant to meet were:

- Support for network communication, concurrency, and parallelization
- Support for excellent building speed
- Support for memory management

### Support for network communication, concurrency, and parallelization

To get the most out of distributed and multi-core machines, Go has excellent support for networked communication, concurrency, and parallelization.
Golang was expected to achieve this target for internal use in Google, and this target is achieved through the concepts of **goroutines**.
Go excels in utilizing multicore and multiprocessor systems better that many other programming languages, which often lack robust support for these architectures.

### Support for excellent building speed

There was a growing concern to improve the building speed (compilation and linking to produce machine code) of C++ projects, which are heavily used in Google's infrastructure.
The language's development was also motivated by cumbersome dependency management in large C++ projects, particularly due to the overhead of "header files."
Go addresses these issues with its package model, offering clean dependency analysis and fast compilation.

Go offers extremely fast compilation times, with its standard library compiling in under 20 seconds and typical projects in just half a second.
This speed surpasses traditional compiled languages like C or Fortran, aligning more with the productivity seen in dynamic languages due to minimal compile / link times.
Moreover, Go maintains execution speed on par with C / C++, combining rapid development cycles with efficient execution.

### Support for memory management

Because memory problems (e.g., memory leaks) are a long-time problem of C++, Go's designers decided that memory management should not be the responsibility of a developer.
So although Go executives native code, it runs in a small runtime, which takes care of an efficient and fast garbage collection.
Go also has a built-in runtime reflection capability.
