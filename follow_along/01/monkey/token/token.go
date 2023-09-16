package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

// LET FUNCTION   IDENTIFIER INTEGER   PLUS EQUALS_TO    COMMA, PLUS  


const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"
  
  // Identifiers + literals
  IDENT = "IDENT"
  INT = "INT"

  // Operators
  ASSIGN = "="
  PLUS = "+"

  // Delimiter
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"


  // Keywords
  FUNCTION = "FUNCTION"
  LET = "LET"
)
