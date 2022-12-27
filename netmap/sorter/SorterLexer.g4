lexer grammar SorterLexer;

CASE : 'CASE';
WHEN : 'WHEN';
END  : 'END';

IDENT             : Nondigit (Digit | Nondigit)* ;
fragment Digit    : [0-9] ;
fragment Nondigit : [a-zA-Z_] ;

NUMBER1 : [1-9] Digit* ;
ZERO    : '0' ;

// Taken from antlr4 json grammar with minor corrections.
// https://github.com/antlr/grammars-v4/blob/master/json/JSON.g4
STRING : '"'  (ESC | SAFECODEPOINTDOUBLE)* '"'
       | '\'' (ESC | SAFECODEPOINTSINGLE)* '\'' ;

fragment ESC : '\\' (['"\\/bfnrt] | UNICODE) ;
fragment UNICODE : 'u' HEX HEX HEX HEX ;
fragment HEX : [0-9a-fA-F] ;
fragment SAFECODEPOINTSINGLE : ~ ['\\\u0000-\u001F] ;
fragment SAFECODEPOINTDOUBLE : ~ ["\\\u0000-\u001F] ;

WS : [ \t\n\r] + -> skip ;


