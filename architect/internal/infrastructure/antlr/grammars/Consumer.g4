grammar Consumer;

import Workflow;

program
    : consumerDecl EOF
    ;

consumerDecl
    : CONSUMES identifier LBRACE action* RBRACE
    ;

// Consumer keywords

CONSUMES : 'consumes';