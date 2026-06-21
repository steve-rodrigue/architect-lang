grammar Endpoint;

import Workflow;

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

httpMethod
    : GET
    | POST
    | PUT
    | PATCH
    | DELETE
    ;

// Endpoint keywords

ENDPOINT : 'endpoint';

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