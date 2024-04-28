parser grammar ExprParser;
options { tokenVocab=ExprLexer; }

program
    : line+ EOF
    ;
    
line: (l_opr | opr | rep | empty) comment? NEWLINE;

arg: (HEX COMMA | REG COMMA)* (REG | HEX);

opr : OPCODE (arg | MARK);

l_opr: MARK COLON opr;

comment: HASHTAG (~NEWLINE)* ;

rep: LPAREN OPCODE INT RPAREN;

empty: ;

// ---------------------------------------------------

// DELETE THIS CONTENT IF YOU PUT COMBINED GRAMMAR IN Parser TAB
lexer grammar ExprLexer;

WS: [ \t\r\f]+ -> skip ;
NEWLINE: '\n';

COLON: ':';
COMMA : ',' ;
LPAREN : '(' ;
RPAREN : ')' ;

INT : [0-9]+ ;

REG: 'R'[0-9]+ | 'AR';
MARK: [a-zA-Z_][a-zA-Z_0-9]* ;
OPCODE: [A-Z]+;

HEX: '0x' [0-9A-F]+ ;
HASHTAG: '#' ;
