package lib

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmpty(t *testing.T) {
	p := NewParser(bytes.NewBufferString(""))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected (")
}
func TestParseNormal(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(* 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "*", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "3"}, l.Children[1])
}

func TestParseAdd(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(+ 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "+", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "3"}, l.Children[1])
}

func TestParseSubtract(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(- 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "-", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "3"}, l.Children[1])
}

func TestParseDivide(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(/ 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "/", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "3"}, l.Children[1])
}

func TestParseManyWhiteSpaces(t *testing.T) {
	p := NewParser(bytes.NewBufferString("  (    *   1    3  )     "))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "*", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "3"}, l.Children[1])
}

func TestParseManyNumbers(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(* 1 3 6 7 8)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "*", l.Operator)
	assert.Len(t, l.Children, 5)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "3"}, l.Children[1])
	assert.Equal(t, &LispStatement{Number: "6"}, l.Children[2])
	assert.Equal(t, &LispStatement{Number: "7"}, l.Children[3])
	assert.Equal(t, &LispStatement{Number: "8"}, l.Children[4])
}

func TestParseSimpleSubStmt(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(+ 1 (- 2 3))"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "+", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, "-", l.Children[1].Operator)
	assert.Len(t, l.Children[1].Children, 2)
	assert.Equal(t, "2", l.Children[1].Children[0].Number)
	assert.Equal(t, "3", l.Children[1].Children[1].Number)
}

func TestParseComplexSubStmt(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(+ 1 (- (* 5 6 4) 3))"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "+", l.Operator)
	assert.Len(t, l.Children, 2)
	assert.Equal(t, &LispStatement{Number: "1"}, l.Children[0])
	assert.Equal(t, "-", l.Children[1].Operator)
	assert.Len(t, l.Children[1].Children, 2)
	assert.Equal(t, "*", l.Children[1].Children[0].Operator)
	assert.Len(t, l.Children[1].Children[0].Children, 3)
	assert.Equal(t, "5", l.Children[1].Children[0].Children[0].Number)
	assert.Equal(t, "6", l.Children[1].Children[0].Children[1].Number)
	assert.Equal(t, "4", l.Children[1].Children[0].Children[2].Number)
	assert.Equal(t, "3", l.Children[1].Children[1].Number)
}

func TestParseNoOperator(t *testing.T) {
	p := NewParser(bytes.NewBufferString("("))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected operator")
}

func TestParseNoNumber(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(*"))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected number")
}

func TestParseNoRP(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(* 1"))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected number or ( or )")
}

func TestParseNoEOF(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(* 1 2) 123"))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected EOF")
}

func TestParseIllegal(t *testing.T) {
	p := NewParser(bytes.NewBufferString("A"))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected (")
}

func TestParseFloat(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(* 1.1 2.2)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "*", l.Operator)
	assert.Equal(t, &LispStatement{Number: "1.1"}, l.Children[0])
	assert.Equal(t, &LispStatement{Number: "2.2"}, l.Children[1])
}

func TestParseSubNoOP(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(+ 1 (123))"))

	l, err := p.Parse()

	assert.Nil(t, l)
	assert.Contains(t, err.Error(), "expected operator")
}
