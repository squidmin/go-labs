# Output Formatting

This module aims to provide users with a comprehensive understanding of the various methods available in Go for formatting and printing output.
Below are some code examples to include in the module, each demonstrating a different aspect of output formatting in Go.

1. Basic Formatting with `fmt.Printf` \
   `fmt.Printf` allows for formatted output according to a format specifier. The following code snippet demonstrates the use of `fmt.Printf` to print a string and an integer.
   ```go
   package main
   
   import "fmt"
   
   func PrintfExample() {
       name := "Test User"
       age := 30
       fmt.Printf("Name: %s, Age: %d\n", name, age)
   }
   
   func main() {
       PrintfExample()
   }
   ```
2. Simple Print with `fmt.Println` \
   `fmt.Println` prints its arguments with spaces between them and a newline at the end, without formatting control.
   ```go
   package main
   
   import "fmt"
   
   func PrintlnExample() {
       name := "Test User"
       age := 30
       fmt.Println("Name:", name, "Age:", age)
   }
   ```
3. String Formatting with `fmt.Sprintf`
   ```go
   package main
   
   import "fmt"
   
   func SprintfExample() string {
       name := "Test User"
       age := 30
       formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
       return formattedString
   }
   
   func main() {
       fmt.Println(SprintfExample())
   }
   ```
4. Writing Formatted Output to Arbitrary Writers with `fmt.Fprintf` \
   `fmt.Fprintf` formats according to a format specifier and writes to the specified writer. \
   It allows for formatted output to be written to any `io.Writer` interface, not just standard output. \
   The following code snippet demonstrates the use of `fmt.Fprintf` to write formatted output to a file. \
   ```go
   package main
   
   import (
       "fmt"
       "os"
   )
   
   func FprintfExample() {
       name := "Test User"
       age := 30
       file, err := os.Create("output.txt")
       if err != nil {
           fmt.Println(err)
           return
       }
       defer func(file *os.File) {
           err := file.Close()
           if err != nil {}
       }(file)
   
       if _, err := fmt.Fprintf(file, "Name: %s, Age: %d\n", name, age); err != nil {
           fmt.Println(err)
           return
       }
   }
   
   func main() {
       FprintfExample()
   }
   ```
5. Error Handling with `fmt.Fprintf` \
   Demonstrates checking for errors when using `fmt.Fprintf`.
   ```go
   package main
   
   import (
       "fmt"
       "os"
   )
   
    func FprintfErrorHandlingExample() {
         name := "Test User"
         age := 30
         file, err := os.Create("output.txt")
         if err != nil {
              fmt.Println(err)
              return
         }
         defer func(file *os.File) {
              err := file.Close()
              if err != nil {}
         }(file)
    
         if _, err := fmt.Fprintf(file, "Name: %s, Age: %d\n", name, age); err != nil {
              fmt.Println(err)
              return
         }
    }
   ```
6. Custom Formatting with `fmt.Formatter` \
    `fmt.Formatter` is an interface that allows for custom formatting of types. \
   The following code snippet demonstrates the use of `fmt.Formatter` to customize the formatting of a custom type.
    ```go
    package main

    import (
        "fmt"
    )

    type Person struct {
        Name string
        Age  int
    }

    func (p Person) Format(f fmt.State, c rune) { // Implementing the fmt.Formatter interface
        switch c {
            case 'v': // Using 'v' as a generic verb that could be customized further
            if f.Flag('+') {
                fmt.Fprintf(f, "%s (%d years old)", p.Name, p.Age)
            } else {
                fmt.Fprintf(f, "%s", p.Name)
            }
        }
    }

    func FormatterExample() {
        person := Person{Name: "Test User", Age: 30}
        fmt.Printf("Name: %v\n", person)           // Default format
        fmt.Printf("Custom: %+v\n", person)        // Custom format triggered by '+'
    }

    func main() {
        FormatterExample()
    }
    ```
7. Custom Formatting with `fmt.Stringer`
   ```go
   package main
   
   import "fmt"
   
   type Person struct {
       Name string
       Age  int
   }
   
   func (p Person) String() string { // Implementing the fmt.Stringer interface
       return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
   }
   
   func StringerExample() {
       person := Person{Name: "Test User", Age: 30}
       fmt.Println(person)
   }
   
   func main() {
       StringerExample()
   }
   ```
