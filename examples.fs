```
$ go build .
$ go install
$ toobeci
Welcome to toobeci

> For now, new words are just symbols
> like foo bar baz
> but we can do some math
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
> Unicode is fun 27700 emit
水
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
> ^D
Goodbye.
```
