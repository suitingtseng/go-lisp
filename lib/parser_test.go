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
	assert.Equal(t, []string{"1", "3"}, l.Numbers)
}

func TestParseAdd(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(+ 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "+", l.Operator)
	assert.Equal(t, []string{"1", "3"}, l.Numbers)
}

func TestParseSubtract(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(- 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "-", l.Operator)
	assert.Equal(t, []string{"1", "3"}, l.Numbers)
}

func TestParseDivide(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(/ 1 3)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "/", l.Operator)
	assert.Equal(t, []string{"1", "3"}, l.Numbers)
}

func TestParseManyWhiteSpaces(t *testing.T) {
	p := NewParser(bytes.NewBufferString("  (    *   1    3  )     "))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "*", l.Operator)
	assert.Equal(t, []string{"1", "3"}, l.Numbers)
}

func TestParseManyNumbers(t *testing.T) {
	p := NewParser(bytes.NewBufferString("(* 1 3 6 7 8)"))

	l, err := p.Parse()

	assert.Nil(t, err)
	assert.Equal(t, "*", l.Operator)
	assert.Equal(t, []string{"1", "3", "6", "7", "8"}, l.Numbers)
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
	assert.Contains(t, err.Error(), "expected )")
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
