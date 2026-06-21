grammar Endpoint;

program
    : endpointDecl EOF
    ;

endpointDecl
    : ENDPOINT IDENT httpMethod STRING LBRACE endpointBody* RBRACE
    ;

endpointBody
    : inputBlock
    | action
    ;

inputBlock
    : INPUT LBRACE inputField* RBRACE
    ;

inputField
    : IDENT typeRef inputSourceRule
    ;

inputSourceRule
    : inputSource
    | LPAREN inputSource (PIPE inputSource)* RPAREN
    ;

inputSource
    : inputSourceKind QUESTION?
    ;

inputSourceKind
    : PATH
    | QUERY
    | BODY
    | HEADER
    | COOKIE
    ;

action
    : fetchAction
    | createAction
    | storeAction
    | updateAction
    | emitAction
    | returnAction
    ;

fetchAction
    : FETCH typedVariable FROM IDENT WHERE condition
    ;

createAction
    : CREATE typedVariable assignmentBlock
    ;

storeAction
    : STORE IDENT TO IDENT
    ;

updateAction
    : UPDATE IDENT assignmentBlock
    ;

emitAction
    : EMIT typedVariable assignmentBlock
    ;

returnAction
    : RETURN expression
    ;

typedVariable
    : IDENT COLON typeRef
    ;

assignmentBlock
    : LBRACE assignment* RBRACE
    ;

assignment
    : IDENT ASSIGN expression
    ;

condition
    : expression comparator expression
    ;

comparator
    : EQ
    | NEQ
    | LTE
    | GTE
    | LT
    | GT
    ;

expression
    : functionCall
    | selector
    | value
    ;

functionCall
    : IDENT LPAREN argumentList? RPAREN
    ;

argumentList
    : expression (COMMA expression)*
    ;

selector
    : identifier (DOT identifier)*
    ;

identifier
    : IDENT
    | INPUT
    ;

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

httpMethod
    : GET
    | POST
    | PUT
    | PATCH
    | DELETE
    ;

// Keywords

ENDPOINT : 'endpoint';
INPUT    : 'input';

FETCH    : 'fetch';
FROM     : 'from';
WHERE    : 'where';
CREATE   : 'create';
STORE    : 'store';
TO       : 'to';
UPDATE   : 'update';
EMIT     : 'emit';
RETURN   : 'return';

PATH     : 'path';
QUERY    : 'query';
BODY     : 'body';
HEADER   : 'header';
COOKIE   : 'cookie';

GET      : 'GET';
POST     : 'POST';
PUT      : 'PUT';
PATCH    : 'PATCH';
DELETE   : 'DELETE';

LIST     : 'List';

TRUE     : 'true';
FALSE    : 'false';

// Symbols

LBRACE   : '{';
RBRACE   : '}';
LPAREN   : '(';
RPAREN   : ')';
LBRACKET : '[';
RBRACKET : ']';

LT       : '<';
GT       : '>';
LTE      : '<=';
GTE      : '>=';

EQ       : '==';
NEQ      : '!=';

ASSIGN   : '=';
COLON    : ':';
DOT      : '.';
COMMA    : ',';
PIPE     : '|';

PLUS     : '+';
STAR     : '*';
QUESTION : '?';

// Literals

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

// Comments / whitespace

LINE_COMMENT
    : '//' ~[\r\n]* -> skip
    ;

BLOCK_COMMENT
    : '/*' .*? '*/' -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;