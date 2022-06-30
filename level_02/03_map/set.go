package _3_map

type Set interface {
	Add(value Value)
	Remove(value Value)
	Size() int
	Contain(value Value) bool
}

type Value interface{}

// 空结构 [不占据内存空间]
type void struct{}

type MySet map[Value]void

func (m MySet) Add(value Value) {
	m[value] = void{}
}

func (m MySet) Remove(value Value) {
	delete(m, value)
}

func (m MySet) Size() int {
	return len(m)
}

func (m MySet) Contain(value Value) bool {
	_, ok := m[value]
	return ok
}

func NewSet(size ...int) MySet {
	var s MySet
	l := len(size)
	switch l {
	case 0:
		s = make(MySet)
	default:
		s = make(MySet, size[0])
	}
	return s
}
