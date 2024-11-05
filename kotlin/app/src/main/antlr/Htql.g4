grammar Htql;

// Lexer rules
SELECT      : 'SELECT';
FROM        : 'FROM';
WHERE       : 'WHERE';
AND         : 'AND';
OR          : 'OR';
IS          : 'IS';
NOT         : 'NOT';
NULL        : 'NULL';
STAR        : '*';
COMMA       : ',';
DOT         : '.';
EQUALS      : '=';
IDENTIFIER  : [a-zA-Z_][a-zA-Z_0-9]*;
STRING      : '\'' (~['\r\n])* '\'';
WS          : [ \t\r\n]+ -> skip;
URL         : ('http' | 'https') '://' [a-zA-Z0-9_.-]+ (':' [0-9]+)? ('/' [a-zA-Z0-9_.-/?&%#=]*)? ;

// Parser rules
query           : selectStmt fromStmt whereStmt? ;
selectStmt      : SELECT (STAR | elementList) ;
elementList     : IDENTIFIER (COMMA IDENTIFIER)* ;
fromStmt        : FROM (IDENTIFIER | URL) ;
whereStmt       : WHERE conditionExpr ;
conditionExpr   : condition (logicalOp condition)* ;
condition       : attributeExpr
                | attributeNullCheck
                | NOT attributeExpr ;

attributeExpr   : 'attributes' DOT IDENTIFIER EQUALS STRING ;
attributeNullCheck
                : 'attributes' IS (NOT)? NULL ;

logicalOp       : AND | OR ;
