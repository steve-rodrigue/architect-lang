grammar Workflow;

import Common;

action
    : fetchAction
    | createAction
    | storeAction
    | updateAction
    | emitAction
    | callAction
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

callAction
    : CALL selector assignmentBlock
    ;

returnAction
    : RETURN expression
    ;

typedVariable
    : identifier COLON typeRef
    ;

assignmentBlock
    : LBRACE assignment* RBRACE
    ;

assignment
    : identifier ASSIGN expression
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

// Workflow keywords

FETCH  : 'fetch';
FROM   : 'from';
WHERE  : 'where';
CREATE : 'create';
STORE  : 'store';
TO     : 'to';
UPDATE : 'update';
EMIT   : 'emit';
CALL   : 'call';
RETURN : 'return';

// Workflow symbols

LPAREN : '(';
RPAREN : ')';

LTE    : '<=';
GTE    : '>=';
EQ     : '==';
NEQ    : '!=';

ASSIGN : '=';
COLON  : ':';
DOT    : '.';