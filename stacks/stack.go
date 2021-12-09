package stacks

type ItemType interface{}

type Stack struct {
	items []ItemType
}

func (s *Stack) Push(t ItemType) []ItemType {

	if s.items == nil {
		s.items = []ItemType{}
	}
	s.items = append(s.items, t)
	return s.items
}

func (s Stack) Pop() {

}
