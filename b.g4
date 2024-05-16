grammar b;

options {
     caseInsensitive = true;
}

prog
    : line* EOF
    ;

line
    : (instruction | lbl | repeat)? EOL
    ;

instruction
    : label? opcode argumentlist?
    ;

repeat
    : '(' opcode TIMES ')'
    ;
    
TIMES
    : 'x' [0-9]+
    ;

lbl
    : label ':'
    ;

argumentlist
    : argument (',' argumentlist)?
    ;

label
    : name
    ;

argument
    : prefix_? (number | name | string_ | '*') (('+' | '-') number)?
    | '(' argument ')'
    ;

prefix_
    : '#'
    ;

string_
    : STRING
    ;

name
    : NAME
    ;

number
    : NUMBER | HEX
    ;

opcode
    : ADD
    | AND
    | BNE
    | BRA
    | DEC
    | INC
    | MOVH
    | MOVL
    | MOVS
    | LDR
    | POP
    | PSH
    | STR
    | SUB
    | XOR
    | NOP
    ;

ADD: 'ADD';

AND: 'AND';

BNE: 'BNE';

BRA: 'BRA';

DEC: 'DEC';

INC: 'INC';

MOVH: 'MOVH';

MOVL: 'MOVL';

MOVS: 'MOVS';

LDR: 'LDR';

POP: 'POP';

PSH: 'PSH';

STR: 'STR';

XOR: 'XOR';

SUB: 'SUB';

NOP: 'NOP';

NAME
    : [A-Z] [A-Z0-9."]*
    ;

NUMBER
    : '$'? [0-9A-F]+
    ;
    
HEX
    : '0x' [0-9A-F]+
    ;

COMMENT
    : '#' ~ [\r\n]* -> skip
    ;

STRING
    : '"' ~ ["]* '"'
    ;

EOL
    : [\r\n]+
    ;

WS
    : [ \t] -> skip
    ;
