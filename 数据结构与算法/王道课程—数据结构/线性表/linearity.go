package linearity

import "sync"

var (
	_defaultSize = 100
)

type Linearity struct {
	cap      int
	size     int
	elements []entry
	mu       *sync.Mutex
}

type entry struct {
	val interface{}
}

func NewLinearity(cap int) *Linearity {
	return &Linearity{
		cap:      cap,
		size:     0,
		elements: make([]entry, cap),
		mu:       new(sync.Mutex),
	}
}

func (l *Linearity) AddElement(val interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.elements = append(l.elements, entry{val: val})
}

func (l *Linearity) Len() int {
	return len(l.elements)
}

func (l *Linearity) Insert(val interface{}, index int) {

}

func (l *Linearity) Acquire(index int) (val interface{}) {

}
