package commands

type StringSearcher[T any] struct {
	start *node[T]
}

type StringSearcherEntry[T any] struct {
	key   string
	value T
}

type node[T any] struct {
	nextEntries map[rune]*node[T]
	values      []T
}

func newNode[T any]() *node[T] {
	n := new(node[T])
	n.nextEntries = make(map[rune]*node[T])
	n.values = make([]T, 0)
	return n
}

func NewStringSearcher[T any]() *StringSearcher[T] {
	ss := new(StringSearcher[T])
	ss.start = newNode[T]()
	return ss
}

func NewStringSearcherWith[T any](entries []StringSearcherEntry[T]) *StringSearcher[T] {
	ss := new(StringSearcher[T])
	ss.start = newNode[T]()
	for _, entry := range entries {
		ss.AddEntry(entry.key, entry.value)
	}
	return ss
}

func (s StringSearcher[T]) AddEntry(str string, value T) {
	var n *node[T] = s.start
	n.values = append(n.values, value)
	for _, ch := range str {
		if n.nextEntries[ch] == nil {
			n.nextEntries[ch] = newNode[T]()
		}
		n = n.nextEntries[ch]
		n.values = append(n.values, value)
	}
}

func (s StringSearcher[T]) GetAvailableValues(str string) []T {
	var n *node[T] = s.start
	for _, ch := range str {
		if n.nextEntries[ch] == nil {
			return []T{}
		}
		n = n.nextEntries[ch]
	}
	return n.values
}
