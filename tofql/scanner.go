package tofql

import (
	"bufio"
	"bytes"
	"io"
)

type Scanner struct {
	r *bufio.Reader
}

const eof = rune(0)

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}
func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isDigit(ch) || ch == '-' {
		s.unread()
		s.scanNumber()
	} else if isIdentChar(ch) {
		s.unread()
		return s.scanIdent()
	}
	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
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

func (s *Scanner) scanNumber() (tok Token, lit string) {
	var buf bytes.Buffer
	ch := s.read()
	if ch == '-' {
		buf.WriteRune(ch)
		ch1 := s.read()
		s.unread()
		if !isDigit(ch1) {
			return ILLEGAL, ch
		}
	} else {
		s.unread()
	}

	buf.WriteString(s.scanDigits())
	if ch = s.read(); ch == '.' {
		buf.WriteRune(ch)

		buf.WriteString(s.scanDigits())
	}

}

// scanDigits consume a contiguous series of digits.
func (s *Scanner) scanDigits() string {
	var buf bytes.Buffer
	for {
		ch, _ := s.read()
		if !isDigit(ch) {
			s.unread()
			break
		}
		_, _ = buf.WriteRune(ch)
	}
	return buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	return IDENT, buf.String()
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }
func isLetter(ch rune) bool     { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }
func isDigit(ch rune) bool      { return (ch >= '0' && ch <= '9') }
func isIdentChar(ch rune) bool {
	return ch != '(' && ch != ')' && ch != '{' && ch != '}' &&
		ch != ',' && ch != '=' && ch != '.' && ch != '\'' && ch != '"' &&
		ch != '\\' && ch != '[' && ch != ']'
}
