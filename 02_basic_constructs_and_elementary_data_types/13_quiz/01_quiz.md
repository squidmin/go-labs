# Quiz: Deduce the Outputs

In this section, your concepts will be evaluated through a quiz.

Choose one correct answer for each question.

### 1. Deduce the output of the following program.

```go
package main

var a = "G"

func main() {
	n()
	
	m()
	n()
}

func n() { print(a) }

func m() {
	a := "0"
	print(a)
}
```

**A**. `0G0`\
**B**. `G0G`\
**C**. `G00`\
**D**. `0GG`

> **Answer**: **B**. `G0G`

### 2. Deduce the output of the following program.

```go
package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	a = "0"
	print(a)
}
```

**A**. `0G0`\
**B**. `G0G`\
**C**. `G00`\
**D**. `0GG`

> **Answer**: **C**. `G00`

### 3. Deduce the output of the following program.

```go
package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "0"
	print(a)
	f2()
}

func f2() {
	print(a)
}
```

**A**. `G0G`\
**B**. `0G0`\
**C**. `0GG`\
**D**. `G00`

> **Answer**: **A**. `G0G`
