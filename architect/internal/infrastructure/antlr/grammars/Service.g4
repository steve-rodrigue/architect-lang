grammar Service;

program
    : serviceDecl EOF
    ;

serviceDecl
    : SERVICE IDENT serviceKind LBRACE serviceBody* RBRACE
    ;

serviceKind
    : IDENT
    ;

serviceBody
    : versionDecl
    | imageDecl
    | exposeDecl
    | dependsOnDecl
    | applicationDecl
    | storeDecl
    | eventDecl
    ;

versionDecl
    : VERSION STRING
    ;

imageDecl
    : IMAGE STRING
    ;

exposeDecl
    : EXPOSES INT
    ;

dependsOnDecl
    : DEPENDS_ON IDENT
    ;

applicationDecl
    : APPLICATION STRING
    ;

storeDecl
    : STORES IDENT
    ;

eventDecl
    : EVENT IDENT
    ;

// Keywords

SERVICE     : 'service';
VERSION     : 'version';
IMAGE       : 'image';
EXPOSES     : 'exposes';
DEPENDS_ON  : 'depends_on';
APPLICATION : 'application';
STORES      : 'stores';
EVENT       : 'event';

// Symbols

LBRACE : '{';
RBRACE : '}';

// Literals

IDENT
    : [a-zA-Z_][a-zA-Z0-9_]*
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