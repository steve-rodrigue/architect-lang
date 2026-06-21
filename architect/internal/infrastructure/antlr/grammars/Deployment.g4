grammar Deployment;

program
    : deploymentDecl EOF
    ;

deploymentDecl
    : DEPLOYMENT IDENT LBRACE deploymentBody* RBRACE
    ;

deploymentBody
    : vendorDecl
    | inventoryDecl
    | serviceDeploymentDecl
    ;

vendorDecl
    : VENDOR IDENT
    ;

inventoryDecl
    : INVENTORY STRING
    ;

serviceDeploymentDecl
    : SERVICE IDENT LBRACE serviceDeploymentBody* RBRACE
    ;

serviceDeploymentBody
    : replicasDecl
    | domainDecl
    | hostDecl
    | volumeDecl
    | backupDecl
    ;

replicasDecl
    : REPLICAS INT
    ;

domainDecl
    : DOMAIN STRING
    ;

hostDecl
    : HOST STRING
    ;

volumeDecl
    : VOLUME STRING
    ;

backupDecl
    : BACKUP booleanValue
    ;

booleanValue
    : TRUE
    | FALSE
    ;

// Keywords

DEPLOYMENT : 'deployment';
VENDOR     : 'vendor';
INVENTORY  : 'inventory';
SERVICE    : 'service';

REPLICAS   : 'replicas';
DOMAIN     : 'domain';
HOST       : 'host';
VOLUME     : 'volume';
BACKUP     : 'backup';

TRUE       : 'true';
FALSE      : 'false';

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