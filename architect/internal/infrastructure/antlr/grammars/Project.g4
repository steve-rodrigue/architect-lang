grammar Project;

program
    : projectDecl EOF
    ;

projectDecl
    : PROJECT IDENT LBRACE versionDecl+ RBRACE
    ;

versionDecl
    : VERSION STRING LBRACE versionBody* RBRACE
    ;

versionBody
    : servicesBlock
    | objectsBlock
    | deploymentsBlock
    | nextVersionDecl
    ;

nextVersionDecl
    : NEXT_VERSION STRING LBRACE versionBody* RBRACE
    ;

servicesBlock
    : SERVICES LBRACE fileRef* RBRACE
    ;

objectsBlock
    : OBJECTS LBRACE fileRef* RBRACE
    ;

deploymentsBlock
    : DEPLOYMENTS LBRACE fileRef* RBRACE
    ;

fileRef
    : STRING
    ;

// Keywords

PROJECT      : 'project';
VERSION      : 'version';
NEXT_VERSION : 'next_version';

SERVICES    : 'services';
OBJECTS     : 'objects';
DEPLOYMENTS : 'deployments';

// Symbols

LBRACE : '{';
RBRACE : '}';

// Literals

IDENT
    : [a-zA-Z_][a-zA-Z0-9_]*
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