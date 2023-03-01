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
> 1 2 . .
2
1
> 1 2 swap . .
1
2
> 1 2 3 rot . . .
1
3
2
> 1 2 over . . .
2
1
2
> clr     \ clears the stack
> .s      \ shows the stack
> 1 2 3 .s
	1
	2
	3
> ^D
Goodbye.
```
