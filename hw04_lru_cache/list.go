package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{Value: v, Prev: l.back}

	if l.back != nil {
		l.back.Next = item
	} else {
		l.front = item
	}

	l.back = item

	l.len++

	return item
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{Value: v, Next: l.front}
	if l.front != nil {
		l.front.Prev = item
	} else {
		l.back = item
	}

	l.front = item

	l.len++

	return item
}

func (l *list) MoveToFront(i *ListItem) {
	if l.front == i {
		return
	}

	l.Remove(i)

	item := i
	item.Next = l.front
	item.Prev = nil

	if l.front != nil {
		l.front.Prev = item
	} else {
		l.back = item
	}

	l.front = item

	l.len++
}

func (l *list) Remove(i *ListItem) {
	prev, next := i.Prev, i.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.front = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		l.back = prev
	}

	i.Prev = nil
	i.Next = nil

	l.len--
}

func NewList() List {
	return new(list)
}
