package doublelist

type List struct {
	root Element
	len  int
}

type Element struct {
	next, prev *Element
	val        interface{}
	list       *List
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) PushBack(val interface{}) *Element {
	l.lazyInit()
	return l.insertValue(val, l.root.prev)
}

//func (l *List) PushFront(val interface{}) *Element {
//
//}

func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{val: v}, at)
}

// insert. insert e after at
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}
