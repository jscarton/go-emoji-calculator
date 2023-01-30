# Go Emoji Calculator

This code implements a calculator able to parse expressions containing  Emoji unicode characters like :100: :keycap_ten: :heavy_plus_sign: or :8ball:


the `/parser` package was generated using [ANTLR4](https://github.com/antlr/antlr4). 

the source code for the Grammar used to generate the parser code is on file `/parser/EmojiCalculator.g4`

```
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
```

the `/calculator` package implements the Listener interface of ANTLR4 parser and the logic to perform the calculations

the `/translator` package implements the function to translate expressions containing emojis into regular ascii expressions and numbers into string containing emojis

the `/challenge` folder contains the challenge which inspired this code and some string examples

# requisites

* Go 1.19

* Install required ANTLR4 runtime go package

```
go get github.com/antlr/antlr4/runtime/Go/antlr/v4
```

# building

On Unix/Linux/Macos:

```
bash build.sh
```

On Windows:

```
build.bat
```

# running

On Unix/Linux/Macos:

```
echo -n "4️2️+🎱+25✖️2" | ./emoji-calculator
```

On Windows:

```
echo "4️2️+🎱+25✖️2" | ./emoji-calculator
```

# running tests

On Unix/Linux/Macos:

```
bash run_tests.sh
```

On Windows:

```
run_tests.bat
```



