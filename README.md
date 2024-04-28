Basic assembly like language compiler for testing my computer organization assignment. ANTLR4 used as parser generator and programmed in go. b.g4 file mostly based on asm6502 example code. z.g4 was my earlier failed attemped.

Install antlr-tools from pip and then compile with this:
```
antlr4 -Dlanguage=Go -o parser b.g4
```

Not tested and probabily produces wrong binaries. Use at your own risk. http://lab.antlr.org/ is a great source to visualize. 
