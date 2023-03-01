# toobeci

<img src="/toobeci.jpg" width="500">

![build](https://github.com/eigenhombre/toobeci/actions/workflows/build.yml/badge.svg)

A simple [Forth](https://en.wikipedia.org/wiki/Forth_(programming_language)) interpreter written in Go.

Mostly just to help me learn both languages better.

<!-- The following examples are autogenerated, do not change by hand! -->
<!-- BEGIN EXAMPLES -->
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
<!-- END EXAMPLES -->

These examples were autogenerated from the tests in `main_test.go`.

Supported operators:

```
+ * / - rot dup drop swap over . .s emit clr
```

# Up Next

- `if`
- `?DO ... LOOP`
- `begin`
- `@`
- `!`
- `s" ...`
- `:`

...
