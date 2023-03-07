```
$ go build .
$ go install
$ toobeci
Welcome to toobeci

> \ comments are ignored
> \ . prints the 'top of the stack':
> 1 .
1
> \ you can do math ...
> 1 2 +
> \ and then show the result:
> .
3
> 10 10 / .
1
> 10 dup dup * * .
1000
> 2 3 drop .
2
> 42 emit
*
> \ The `emit` operator emits unicode characters:
> 27700 emit
水
> clr      \ clears the stack
> .s       \ shows the stack
> 1 2 3 .s
	3
	2
	1
> swap .s  \ swap top two items
	2
	3
	1
> rot .s   \ rotate items
	1
	2
	3
> over .s  \ copy & promote 2nd item
	2
	1
	2
	3
> \ Some boolean logic:
> 1 1 and .
1
> 1 0 and .
0
> 1 1 or .
1
> 1 0 or .
1
> \ Default 'true' value is -1 (0b1111...):
> 1 1 = .
-1
> 1 0 = .
0
> 3 not .
0
> 0 not .
-1
> dakine
unknown word: dakine
> ^D
Goodbye.
```
