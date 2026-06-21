grammar Object;

import Common;

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
    : IDENT typeRef defaultValue? fieldModifier*
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

// Keywords

OBJECT       : 'object';
HISTORY_OF   : 'history_of';
PRIMARY      : 'primary';
UNIQUE       : 'unique';
RENAMED_FROM : 'renamed_from';
DEPRECATED   : 'deprecated';

// Symbols

LBRACE : '{';
RBRACE : '}';

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