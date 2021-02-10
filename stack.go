package goforth

type data []string

type stack struct {
	data
}

func newStack() *stack {
	return &stack{
		data: []string{},
	}
}

func (s *stack) push(val string) {
	s.data = append(s.data, val)
}

func (s *stack) pop() string {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}
