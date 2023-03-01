```
$ go build .
$ go install
$ toobeci
Welcome to toobeci

> \ comments are ignored
> 1 2 +
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
> \ Unrecognized symbols are just strings for now.
> \ But the `emit` operator emits unicode characters:
> Unicode is fun 27700 emit
水
> .
fun
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
> ^D
Goodbye.
```
