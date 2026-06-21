grammar Common;

typeRef
    : typeName numberConstraint? optionalMarker?
    | LIST LT typeName GT optionalMarker?
    ;

typeName
    : IDENT
    ;

numberConstraint
    : PLUS
    | STAR
    | LBRACKET numberValue? COMMA numberValue? RBRACKET
    ;

optionalMarker
    : QUESTION
    ;

value
    : STRING
    | INT
    | FLOAT
    | TRUE
    | FALSE
    ;

numberValue
    : INT
    | FLOAT
    ;

identifier
    : IDENT
    | INPUT
    | EVENT
    | RESULT
    ;

LIST   : 'List';
INPUT  : 'input';
EVENT  : 'event';
RESULT : 'result';

TRUE   : 'true';
FALSE  : 'false';

LBRACE   : '{';
RBRACE   : '}';
LPAREN   : '(';
RPAREN   : ')';
LBRACKET : '[';
RBRACKET : ']';

LT       : '<';
GT       : '>';

COMMA    : ',';
PIPE     : '|';

PLUS     : '+';
STAR     : '*';
QUESTION : '?';

IDENT
    : [a-zA-Z_][a-zA-Z0-9_]*
    ;

FLOAT
    : [0-9]+ '.' [0-9]+
    ;

INT
    : [0-9]+
    ;

STRING
    : '"' (~["\\] | '\\' .)* '"'
    ;

LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;