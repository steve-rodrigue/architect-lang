grammar Application;

program
    : applicationDecl EOF
    ;

applicationDecl
    : APPLICATION IDENT LBRACE applicationBody* RBRACE
    ;

applicationBody
    : portDecl
    | endpointsBlock
    | consumersBlock
    ;

portDecl
    : EMITS IDENT ON INT
    | LISTENS IDENT ON INT
    ;

endpointsBlock
    : ENDPOINTS LBRACE fileRef* RBRACE
    ;

consumersBlock
    : CONSUMERS LBRACE fileRef* RBRACE
    ;

fileRef
    : STRING
    ;

APPLICATION : 'application';

EMITS   : 'emits';
LISTENS : 'listens';
ON      : 'on';

ENDPOINTS : 'endpoints';
CONSUMERS : 'consumers';

LBRACE : '{';
RBRACE : '}';

IDENT
    : [a-zA-Z_][a-zA-Z0-9_]*
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