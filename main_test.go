package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/suitingtseng/go-lisp/lib"
)

func TestEvalAdd(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "+",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
			&lib.LispStatement{Number: "2"},
			&lib.LispStatement{Number: "4"},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, 7.0, res)
}

func TestEvalSubtract(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "-",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
			&lib.LispStatement{Number: "2"},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, -1.0, res)
}

func TestEvalMultiply(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "*",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
			&lib.LispStatement{Number: "2"},
			&lib.LispStatement{Number: "4"},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, 8.0, res)
}

func TestEvalDivide(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "/",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
			&lib.LispStatement{Number: "2"},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.Equal(t, 0.5, res)
}

func TestEvalSubtractArguments(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "-",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
		},
	}

	_, err := eval(l)

	assert.Equal(t, err.Error(), "substract only accept 2 arguments")
}

func TestEvalDivideArguments(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "/",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
		},
	}

	_, err := eval(l)
	assert.Equal(t, err.Error(), "division only accept 2 arguments")
}

func TestEvalDivideByZero(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "/",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
			&lib.LispStatement{Number: "0"},
		},
	}

	_, err := eval(l)
	assert.Equal(t, err.Error(), "divided by zero")
}

func TestEvalMultiplyFloat(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "*",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1.1"},
			&lib.LispStatement{Number: "2.2"},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.InEpsilon(t, 2.42, res, 0.0001)
}

func TestEvalSubStatement(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "*",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "1"},
			&lib.LispStatement{
				Operator: "+",
				Children: []*lib.LispStatement{
					&lib.LispStatement{Number: "2"},
					&lib.LispStatement{Number: "3"},
				},
			},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.InEpsilon(t, 5, res, 0.0001)
}

func TestEvalComplexSubStatement(t *testing.T) {
	l := &lib.LispStatement{
		Operator: "*",
		Children: []*lib.LispStatement{
			&lib.LispStatement{Number: "5"},
			&lib.LispStatement{
				Operator: "+",
				Children: []*lib.LispStatement{
					&lib.LispStatement{
						Operator: "/",
						Children: []*lib.LispStatement{
							&lib.LispStatement{Number: "2"},
							&lib.LispStatement{Number: "3"},
						},
					},
					&lib.LispStatement{Number: "3"},
				},
			},
		},
	}

	res, err := eval(l)

	assert.Nil(t, err)
	assert.InEpsilon(t, 18.333333, res, 0.0001)
}
