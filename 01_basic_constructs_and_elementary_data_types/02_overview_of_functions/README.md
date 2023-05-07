# Overview of functions


### Naming things in Go

Clean readable code and simplicity are major goals of Go development.
As such, the names of things in Go should be short, concise, and evocative.
Long names with mixed caps and underscores (which are often seen in Java and Python code) sometimes hinder readability.

Names should not contain an indication of the package.

A method or function which returns an object is named as a noun; no `get...` is needed.
To change an object, use `set...`.

If necessary, Go uses`MixedCaps` or `mixedCaps` rather than underscores to write multi-word names.


### Function syntax

```go
func exampleFunction(param_1 type_1, ..., param_n type_n) (ret_1 type_1, ..., ret_n type_n) {
    ...
}
```
