package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	character    byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readCharacter()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	l.skipWhitespace()

	switch l.character {
	case '=':
		if l.peekChar() == '=' {
			character := l.character
			l.readCharacter()
			literal := string(character) + string(l.character)
			nextToken = token.Token{Type: token.EQ, Literal: literal}
		} else {
			nextToken = newToken(token.ASSIGN, l.character)
		}
	case '!':
		if l.peekChar() == '=' {
			character := l.character
			l.readCharacter()
			literal := string(character) + string(l.character)
			nextToken = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			nextToken = newToken(token.BANG, l.character)
		}
	case '+':
		nextToken = newToken(token.PLUS, l.character)
	case '-':
		nextToken = newToken(token.MINUS, l.character)
	case '/':
		nextToken = newToken(token.SLASH, l.character)
	case '*':
		nextToken = newToken(token.ASTERISK, l.character)
	case '<':
		nextToken = newToken(token.LT, l.character)
	case '>':
		nextToken = newToken(token.GT, l.character)
	case ',':
		nextToken = newToken(token.COMMA, l.character)
	case ';':
		nextToken = newToken(token.SEMICOLON, l.character)
	case '(':
		nextToken = newToken(token.LPAREN, l.character)
	case ')':
		nextToken = newToken(token.RPAREN, l.character)
	case '{':
		nextToken = newToken(token.LBRACE, l.character)
	case '}':
		nextToken = newToken(token.RBRACE, l.character)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	default:
		if isLetter(l.character) {
			nextToken.Literal = l.readIdentifier()
			nextToken.Type = token.LookupIdent(nextToken.Literal)

			return nextToken
		} else if isDigit(l.character) {
			nextToken.Type = token.INT
			nextToken.Literal = l.readNumber()

			return nextToken
		} else {
			nextToken = newToken(token.ILLEGAL, l.character)
		}
	}

	l.readCharacter()

	return nextToken
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(character),
	}
}

func (l *Lexer) skipWhitespace() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readCharacter()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.character) {
		l.readCharacter()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.character) {
		l.readCharacter()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readCharacter() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		l.character = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
