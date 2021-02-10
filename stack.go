package goforth

type stack []string

type stackWrapper struct {
	stack
}

func newStack() *stackWrapper {
	return &stackWrapper{
		stack: []string{},
	}
}

func (s *stackWrapper) push(val string) {
	s.stack = append(s.stack, val)
}

func (s *stackWrapper) pop() string {
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}
