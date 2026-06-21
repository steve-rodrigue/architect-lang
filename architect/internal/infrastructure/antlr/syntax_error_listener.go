package antlr

import (
	"fmt"

	antlr4 "github.com/antlr4-go/antlr/v4"
)

type syntaxErrorListener struct {
	*antlr4.DefaultErrorListener
	err error
}

func (l *syntaxErrorListener) SyntaxError(
	recognizer antlr4.Recognizer,
	offendingSymbol interface{},
	line int,
	column int,
	msg string,
	e antlr4.RecognitionException,
) {
	if l.err == nil {
		l.err = fmt.Errorf("syntax error at line %d:%d: %s", line, column, msg)
	}
}
