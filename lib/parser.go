package lib

import (
	"bytes"
	"fmt"
	"io"
)

type LispStatement struct {
	Number   string
	Operator string
	Children []*LispStatement
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

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != LPARENTHESIS && tok != NUMBER {
		return nil, fmt.Errorf("found %q, expected ( or number", lit)
	} else if tok == NUMBER {
		l.Number = lit
		return l, nil
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != OPERATOR {
		return nil, fmt.Errorf("found %q, expected operator", lit)
	} else {
		l.Operator = lit
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok == NUMBER || tok == LPARENTHESIS {
		p.unscan()

		for {
			tok, lit := p.scanIgnoreWhiteSpace()
			if tok == RPARENTHESIS {
				p.unscan()
				break
			} else if tok == NUMBER {
				buf := bytes.NewBufferString(lit)
				cp := NewParser(buf)
				cl, err := cp.parse()
				if err != nil {
					return nil, err
				}
				l.Children = append(l.Children, cl)
			} else if tok == LPARENTHESIS {
				stack := 1
				buf := bytes.NewBufferString(lit)
				for {
					tok, lit := p.scanIgnoreWhiteSpace()
					if tok == RPARENTHESIS {
						buf.WriteString(lit)
						stack -= 1
						if stack == 0 {
							break
						}
					} else if tok == LPARENTHESIS {
						buf.WriteString(lit)
						stack += 1
					} else if tok == EOF {
						return nil, fmt.Errorf("found %q, expected number or )", lit)
					} else {
						buf.WriteString(lit)
						buf.WriteString(" ")
					}
				}
				cp := NewParser(buf)
				cl, err := cp.parse()
				if err != nil {
					return nil, err
				}
				l.Children = append(l.Children, cl)
			} else {
				return nil, fmt.Errorf("found %q, expected number or ( or )", lit)
			}
		}
	} else {
		return nil, fmt.Errorf("found %q, expected number or (", lit)
	}

	if tok, lit := p.scanIgnoreWhiteSpace(); tok != RPARENTHESIS {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}

	return l, nil
}
