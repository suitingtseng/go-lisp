package lib

import (
	"bufio"
	"bytes"
	"io"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	LPARENTHESIS
	RPARENTHESIS

	NUMBER

	OPERATOR
)

type Operator int

const (
	ADD Operator = iota
	SUBSTRACT
	MULTIPLY
	DIVIDE
)

var eof = rune(0)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isNumber(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isLPARENTHESIS(ch rune) bool {
	return ch == '('
}

func isRPARENTHESIS(ch rune) bool {
	return ch == ')'
}

func isEOF(ch rune) bool {
	return ch == eof
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}

func (s *Scanner) Scan() (Token, string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isNumber(ch) {
		s.unread()
		return s.scanNumber()
	} else if isOperator(ch) {
		return OPERATOR, string(ch)
	} else if isLPARENTHESIS(ch) {
		return LPARENTHESIS, string(ch)
	} else if isRPARENTHESIS(ch) {
		return RPARENTHESIS, string(ch)
	} else if isEOF(ch) {
		return EOF, string(ch)
	} else {
		return ILLEGAL, string(ch)
	}
}

func (s *Scanner) scanWhitespace() (Token, string) {
	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanNumber() (Token, string) {
	var buf bytes.Buffer

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isNumber(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return NUMBER, buf.String()
}
