# GoForth

After reading about the [Forth](https://en.wikipedia.org/wiki/Forth_(programming_language)) programming language and its relationship to bitcoin script I decided to have a bash and writing an example in go, because, why not!

This is a stack based (LIFO) [reverse polish](https://en.wikipedia.org/wiki/Reverse_Polish_notation) language.

It has a limited set of simple operators as this is just a toy and learning experience at the moment, these are listed below:

| Operator 	| Description                                                                                	|
|----------	|--------------------------------------------------------------------------------------------	|
| ADD      	| Adds two numbers and pushes the result to the stack                                        	|
| SUB      	| Takes a number from another and pushes the result to the stack                             	|
| EQ       	| Evaluates if the preceeding two values are equal, it will push 1 if TRUE or 0 if FALSE     	|
| NE       	| Evaluates if the preceeding two values are not equal, it will push 1 if TRUE or 0 if FALSE 	|
| HASH256  	| Generates a SHA256 from the preceeding value and pushes to the stack                       	|

More can be easily added by adding to the ops.go file - feel free to create an issue and PR if you want some more for some reason.

## Useage

It's a very simple public interface, just one function that takes your expression string and returns the last remaining item on the stack after evaluating every word in the expression.

To add two numbers for example, you would supply this string `2 3 ADD` which would evaluate to 5.

In terms of this library you would do:

```go
result := goforth.Run("1 10 ADD 11 EQ HASH256")
```

## Contributing

If for whatever reason you have a use for this and want to add more operators create an issue and PR and add them to the `ops.go` file.

Same goes if you spot an error, which is highly likely as I just threw this together very quickly.
