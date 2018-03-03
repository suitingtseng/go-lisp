# go-lisp

Writing a lisp interpreter in Golang as an excercise.

### Usage

```shell
$ go-lisp --help
usage: go-lisp <filename>
```

### Examples

```shell
$ go-lisp <(echo '(* 1 2 3.5)')
Result: 7.0000
```

# Project status

It's in a very early stage. 

- [x] arithmetics with int
- [x] arithmetics with float
- [x] sub-statement

# References

- [Handwritten Parsers & Lexers in Go](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)
