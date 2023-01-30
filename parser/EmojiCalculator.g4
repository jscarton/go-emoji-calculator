grammar EmojiCalculator;

// Possible Tokens
MULTIPLICATION: '*' | 'x' | 'times';
DIVISION: '/' | '%' | 'divided by';
ADDITION: '+' | 'plus';
SUBSTRACTION: '-' | 'minus';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

expression
   : expression op=(MULTIPLICATION | DIVISION) expression # MulDiv
   | expression op=(ADDITION | SUBSTRACTION) expression # AddSub
   | NUMBER                             # Number
   ;