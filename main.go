package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/suitingtseng/go-lisp/lib"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage()
		os.Exit(0)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	p := lib.NewParser(f)
	listStatement, err := p.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	res, err := eval(listStatement)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
	fmt.Printf("Result: %.4f\n", res)
}

func usage() {
	fmt.Println("usage: go-lisp <filename>")
}

func eval(l *lib.LispStatement) (float64, error) {
	op := l.Operator
	num := l.Number
	children := l.Children
	if num != "" {
		val, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number %q", err)
		}
		return val, nil
	}

	switch op {
	case "+":
		var sum float64
		for _, c := range children {
			val, err := eval(c)
			if err != nil {
				return 0, err
			}
			sum += val
		}
		return sum, nil
	case "*":
		var accu float64
		accu = 1
		for _, c := range children {
			val, err := eval(c)
			if err != nil {
				return 0, err
			}
			accu *= val
		}
		return accu, nil
	case "-":
		if len(children) != 2 {
			return 0, errors.New("substract only accept 2 arguments")
		}
		num1, err := eval(children[0])
		if err != nil {
			return 0, err
		}
		num2, err := eval(children[1])
		if err != nil {
			return 0, err
		}
		return num1 - num2, nil
	case "/":
		if len(children) != 2 {
			return 0, errors.New("division only accept 2 arguments")
		}
		num1, err := eval(children[0])
		if err != nil {
			return 0, err
		}
		num2, err := eval(children[1])
		if err != nil {
			return 0, err
		}
		if num2 == 0 {
			return 0, errors.New("divided by zero")
		}
		return num1 / num2, nil
	}
	return 0, nil
}
