Emoji Calculator
================

Emoji are everywhere these days, and at least for humans,
it can be very different to distinguish between emoji
symbols and their ASCII cousins. In fact this has been
the basis of a number of DNS based attacks where people
register two almost identical domain names one which
contains a unicode (so technically different) version
of a domain name, like MojoHost.com vs MojoHÐ¾st.com (those are
not the same string, one of them has an ASCII a, the other
one has a cyrilic small letter Ð¾ (`U+043E`).

For this puzzle, we'll consider a happier example, the
fact that calculators, not being humans, don't look at
the shape of the symbols but rather at the bits. As a
result they get very confused when passed in Emoji numbers
like 1ï¸, or ð or even ð¯.

Write a simple arithmetic calculator capable of computing
statements involving addition, subtraction, 
multiplication and division. The calculator must accept
input as either ASCII character, words or emoji (see the
alphabet below for a complete input dictionary). The output
should be a minimal number of emojii symbols necessary
to express the answer or ð¤· in the event of an error.

Your program should be robust to input oddities, for
example extra space surrounding operations, and should
be conservative (and consistent in it's output).

Your program should read in an expression from `stdin`
and print your answer on `stdout`.

Example Data
------------

The file `examples.txt` contains a list of inputs, followed
by their expected output on the next line. So, for example,
the first two lines are:

> 4ï¸2ï¸+ð±+ 25âï¸2
>
> ð¯

The first line is an expression and the second is the
expected output.

Input Alphabet
--------------

Your program should accept the follow input, mixed in any
combination:

*Numbers*: ASCII digits 0-9, Emoji Digits 0ï¸-9ï¸, ð, ð¯ and ð±.

*Addition*: ASCII +, Emoji â or the word `plus`.

*Subtraction*: ASCII -, Emoji â or the word `minus`.

*Multiplication*: ASCII *, ASCII x, Emoji âï¸ or the word `times`.

*Division*: ASCII /, ASCII %, Emoji â or the words `divided by`.

Output Alphabet
---------------

You program should output either a number, consisting of
Emoji Digits 0ï¸-9ï¸, ð, and ð¯, or the ð¤· emoji if an
error occurred and no answer could be provided. Each answer
should be terminated with a single newline character (`\n`).
Try to express the output with the minimal number of symbols.
In the event an error occurred, please also print a
human readable explanation for the error to `stdout` (in
ASCII).

Important Notes
---------------

Recall that the order of operations are that multiplication
and division take precedence over addition and subtraction,
but as between multiplication and division expressions are
evaluated left to right.

Also keep in mind that its valid to have space between a
number and an operator but not between digits of the same
number.

It's perfectly valid to just input a valid number and have
the calculator produce that number as output.
