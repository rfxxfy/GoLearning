type Stack struct {
	data []int
}

func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

type Queue struct {
	leftStack  Stack
	rightStack Stack
}

func (q *Queue) Push(v int) {
	q.leftStack.Push(v)
}

func (q *Queue) Pop() (int, bool) {
	if q.rightStack.IsEmpty() {
		if q.leftStack.IsEmpty() {
			return 0, false
		}
		for !q.leftStack.IsEmpty() {
			v, _ := q.leftStack.Pop()
			q.rightStack.Push(v)
		}
	}
	return q.rightStack.Pop()
}

func (q *Queue) IsEmpty() bool {
	return q.leftStack.IsEmpty() && q.rightStack.IsEmpty()
}

func (q *Queue) PrintQueue() {
	tempStack := Stack{}
	for !q.rightStack.IsEmpty() {
		v, _ := q.rightStack.Pop()
		tempStack.Push(v)
	}
	for !q.leftStack.IsEmpty() {
		v, _ := q.leftStack.Pop()
		tempStack.Push(v)
	}

	for !tempStack.IsEmpty() {
		v, _ := tempStack.Pop()
		fmt.Println(v)
		q.rightStack.Push(v)
	}
}
