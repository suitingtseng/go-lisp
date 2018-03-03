package lib

import (
	"fmt"
	"io"
)

type LispStatement struct {
	Operator string
	Numbers  []string
}

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) scan() (tok Token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) unscan() {
	p.buf.n = 1
}

func (p *Parser) scanIgnoreWhiteSpace() (Token, string) {
	tok, lit := p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return tok, lit
}

func (p *Parser) Parse() (*LispStatement, error) {
	l, err := p.parse()
	if err != nil {
		return nil, err
	}

	// parse the final EOF
	if tok, lit := p.scanIgnoreWhiteSpace(); tok != EOF {
		return nil, fmt.Errorf("found %q, expected EOF", lit)
	}
	return l, nil
}

func (p *Parser) parse() (*LispStatement, error) {
	l := &LispStatement{}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != LPARENTHESIS {
		return nil, fmt.Errorf("found %q, expected (", lit)
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != OPERATOR {
		return nil, fmt.Errorf("found %q, expected operator", lit)
	} else {
		l.Operator = lit
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != NUMBER {
		return nil, fmt.Errorf("found %q, expected number", lit)
	} else {
		p.unscan()

		for {
			if tok, lit := p.scanIgnoreWhiteSpace(); tok != NUMBER {
				p.unscan()
				break
			} else {
				l.Numbers = append(l.Numbers, lit)
			}
		}
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != RPARENTHESIS {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}

	return l, nil
}
