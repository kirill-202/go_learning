package main


type List[T comparable] struct {
	next *List[T]
	val T
}

func (l *List[T]) Add(value T) {
	current := l 
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
}

func (l *List[T]) Find(value T) *List[T] {
    current := l
    for current != nil {
        if current.val == value {
            return current
        }
        current = current.next
    }
    return nil
}