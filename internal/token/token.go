package token

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Line     int
	Column   int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"


	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT  = "FLOAT"
	STRING = "STRING"


	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG   = "!"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"

	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="
	LTE = "<="
	GTE = ">="


	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"


	NUMERO   = "numero"
	TEXTO    = "texto"
	BOOLEANO = "booleano"
	LISTA    = "lista"
	MAPA     = "mapa"
	FUNÇÃO   = "função"
	FUNCAO   = "funcao"
	RETORNE  = "retorne"
	SE       = "se"
	SENAO    = "senao"
	SENÃO    = "senão"
	PARA     = "para"
	ATÉ      = "até"
	ATE      = "ate"
	FIM      = "fim"
	E        = "e"
	OU       = "ou"
	NÃO      = "não"
	NAO      = "nao"
	VERDADEIRO = "verdadeiro"
	FALSO      = "falso"
	IMPRIMIR   = "imprimir"
	INCLUDE    = "include"
	INSTALL    = "install"
	COMMENT    = "COMMENT"
	NEWLINE    = "NEWLINE"
)


var keywords = map[string]TokenType{
	"numero":     NUMERO,
	"texto":      TEXTO,
	"booleano":   BOOLEANO,
	"lista":      LISTA,
	"mapa":       MAPA,
	"função":     FUNÇÃO,
	"funcao":     FUNCAO,
	"retorne":    RETORNE,
	"se":         SE,
	"senao":      SENAO,
	"senão":      SENÃO,
	"para":       PARA,
	"até":        ATÉ,
	"ate":        ATE,
	"fim":        FIM,
	"e":          E,
	"ou":         OU,
	"não":        NÃO,
	"nao":        NAO,
	"verdadeiro": VERDADEIRO,
	"falso":      FALSO,
	"imprimir":   IMPRIMIR,
	"include":    INCLUDE,
	"install":    INSTALL,
}


func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
