package hw04lrucache

type ListItem struct {
	Next  *ListItem
	Prev  *ListItem
	Value interface{}
}

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	count     int
	backItem  *ListItem
	frontItem *ListItem
}

func NewList() List { return new(list) }

func (l list) Len() int {
	return l.count
}

func (l list) Front() *ListItem {
	return l.frontItem
}

func (l list) Back() *ListItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{Value: v}

	if l.count == 0 {
		// список пуст, newListItem также будет и последним элементом
		l.backItem = newListItem
	} else {
		// список не пуст, правим ссылки
		newListItem.Next = l.frontItem
		l.frontItem.Prev = newListItem
	}

	// заменяем первый элемент
	l.frontItem = newListItem
	l.count++

	return newListItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := &ListItem{Value: v}

	if l.count == 0 {
		// список пуст, newListItem также будет и первым элеметом
		l.frontItem = newListItem
	} else {
		// список не пуст, правим ссылки
		newListItem.Prev = l.backItem
		l.backItem.Next = newListItem
	}

	// заменяем первый элемент
	l.backItem = newListItem
	l.count++

	return newListItem
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Prev == nil:
		// первый элемент
		i.Next.Prev = nil
		l.frontItem = i
	case i.Next == nil:
		// последний элемент
		i.Prev.Next = nil
		l.frontItem = i
	default:
		// не крайний элемент
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i.Prev == nil:
		// первый элемент
		return
	case i.Next == nil:
		// последний элемент
		i.Prev.Next = nil
		l.backItem = i.Prev
	default:
		// не крайний элемент
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	// предыдущий первый элемент, меняем ссылки
	oldFirstItem := l.frontItem
	oldFirstItem.Prev = i
	i.Next = oldFirstItem

	// новый первый элемент
	l.frontItem = i
}
