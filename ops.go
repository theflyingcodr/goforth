package goforth

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
)

var ops = map[string]func(s *stackWrapper){
	"ADD":     add,
	"SUB":     sub,
	"EQ":      equal,
	"NE":      notequal,
	"HASH256": hash256,
}

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

func equal(stack *stackWrapper) {
	if stack.pop() == stack.pop() {
		stack.push("1")
		return
	}
	stack.push("0")
}

func notequal(stack *stackWrapper) {
	if stack.pop() != stack.pop() {
		stack.push("1")
		return
	}
	stack.push("0")
}

func hash256(stack *stackWrapper) {
	s := stack.pop()
	h := sha256.Sum256([]byte(s))
	stack.push(fmt.Sprintf("%x", h))
}
