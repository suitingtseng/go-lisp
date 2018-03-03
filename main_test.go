package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suitingtseng/go-lisp/lib"
)

func TestEvalAdd(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "+",
		Numbers:  []string{"1", "2", "4"},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, 7.0, res)
}

func TestEvalSubtract(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "-",
		Numbers:  []string{"1", "2"},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, -1.0, res)
}

func TestEvalMultiply(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "*",
		Numbers:  []string{"1", "2", "4"},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, 8.0, res)
}

func TestEvalDivide(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "/",
		Numbers:  []string{"1", "2"},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, 0.5, res)
}

func TestEvalSubtractArguments(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "-",
		Numbers:  []string{"1"},
	}

	_, err := eval(l)

	assert.Equal(t, err.Error(), "substract only accept 2 arguments")
}

func TestEvalDivideArguments(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "/",
		Numbers:  []string{"1"},
	}

	_, err := eval(l)
	assert.Equal(t, err.Error(), "division only accept 2 arguments")
}

func TestEvalDivideByZero(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "/",
		Numbers:  []string{"1", "0"},
	}

	_, err := eval(l)
	assert.Equal(t, err.Error(), "divided by zero")
}

func TestEvalMultiplyFloat(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "*",
		Numbers:  []string{"1.1", "2.2"},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.InEpsilon(t, 2.42, res, 0.0001)
}
