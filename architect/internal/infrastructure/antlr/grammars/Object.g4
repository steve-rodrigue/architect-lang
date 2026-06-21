grammar Object;

program
    : objectDecl EOF
    ;

objectDecl
    : OBJECT IDENT objectModifier* LBRACE fieldDecl* RBRACE
    ;

objectModifier
    : HISTORY_OF IDENT
    | RENAMED_FROM IDENT
    | DEPRECATED
    ;

fieldDecl
    : IDENT typeRef fieldModifier*
    ;

typeRef
    : typeName numberConstraint? optionalMarker? defaultValue?
    | LIST LT typeName GT optionalMarker? defaultValue?
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

defaultValue
    : value
    ;

fieldModifier
    : PRIMARY
    | UNIQUE
    | RENAMED_FROM IDENT
    | DEPRECATED
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

// Keywords
OBJECT       : 'object';
HISTORY_OF   : 'history_of';
PRIMARY      : 'primary';
UNIQUE       : 'unique';
RENAMED_FROM : 'renamed_from';
DEPRECATED   : 'deprecated';
LIST         : 'List';
TRUE         : 'true';
FALSE        : 'false';

// Symbols
LBRACE   : '{';
RBRACE   : '}';
LBRACKET : '[';
RBRACKET : ']';
LT       : '<';
GT       : '>';
COMMA    : ',';
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