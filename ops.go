package goforth

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
)

// ops is a list of all supported operations and their function for easy execution.
var ops = map[string]func(s *stackWrapper){
	"ADD":     add,
	"SUB":     sub,
	"EQ":      equal,
	"NE":      notequal,
	"HASH256": hash256,
	"DUP":     duplicate,
}

// add will add two numbers together by popping
// them off the stack, converting to int and
// performing the maths.
// It will then push the result onto the stack.
func add(stack *stackWrapper) {
	a := stack.pop()
	b := stack.pop()
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	o := n1 + n2
	stack.push(fmt.Sprintf("%d", o))
}

// sub will minus two numbers together by popping
// them off the stack, converting to int and
// performing the maths.
// This will subtract in the order of the expression, so,
// given 10 3 SUB it will perform 10-3.
// It will then push the result onto the stack.
func sub(stack *stackWrapper) {
	a := stack.pop()
	b := stack.pop()
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	o := n2 - n1
	stack.push(fmt.Sprintf("%d", o))
}

// equal will check the last to items on the stack are equal.
// It will return 1 if TRUE and 0 if FALSE.
// The result will be pushed onto the stack.
func equal(stack *stackWrapper) {
	if stack.pop() == stack.pop() {
		stack.push("1")
		return
	}
	stack.push("0")
}

// notequal will check the last to items on the stack are not equal.
// It will return 1 if TRUE and 0 if FALSE.
// The result will be pushed onto the stack.
func notequal(stack *stackWrapper) {
	if stack.pop() != stack.pop() {
		stack.push("1")
		return
	}
	stack.push("0")
}

// hash256 will hash the current top value in the stack
// using the sha256 algo.
// The result is pushed onto the stack.
func hash256(stack *stackWrapper) {
	s := stack.pop()
	h := sha256.Sum256([]byte(s))
	stack.push(fmt.Sprintf("%x", h))
}

// duplicate will take the top item in the stack and duplicate it.
// It will then push this duplicate value to the top of the stack.
func duplicate(stack *stackWrapper) {
	s := stack.pop()
	stack.push(s)
	stack.push(s)
}
