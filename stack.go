package simplestack

type stack []interface{}

func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Size() int {
	return len(s)
}

func (s stack) Peek() interface{} {
	return s[len(s)-1]
}

func (s *stack) Put(i interface{}) {
	(*s) = append((*s), i)
}

func (s *stack) Pop() interface{} {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}
