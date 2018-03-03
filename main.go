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
	nums := l.Numbers
	switch op {
	case "+":
		var sum float64
		for _, n := range nums {
			val, err := strconv.ParseFloat(n, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number %q", n)
			}
			sum += val
		}
		return sum, nil
	case "*":
		var accu float64
		accu = 1
		for _, n := range nums {
			val, err := strconv.ParseFloat(n, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number %q", n)
			}
			accu *= val
		}
		return accu, nil
	case "-":
		if len(nums) != 2 {
			return 0, errors.New("substract only accept 2 arguments")
		}
		num1, err := strconv.ParseFloat(nums[0], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number %q", nums[0])
		}
		num2, err := strconv.ParseFloat(nums[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number %q", nums[1])
		}
		return num1 - num2, nil
	case "/":
		if len(nums) != 2 {
			return 0, errors.New("division only accept 2 arguments")
		}
		num1, err := strconv.ParseFloat(nums[0], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number %q", nums[0])
		}
		num2, err := strconv.ParseFloat(nums[1], 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number %q", nums[1])
		}
		if num2 == 0 {
			return 0, errors.New("divided by zero")
		}
		return num1 / num2, nil
	}
	return 0, nil
}
