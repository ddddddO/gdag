package gdag

// T is short name of Task func.
func T(text string) *Node {
	return Task(text)
}

// D is short name of Done func.
func D(nodes ...*Node) {
	Done(nodes...)
}

// C is short name of Con func.
func (upstream *Node) C(current *Node) *Node {
	return upstream.Con(current)
}

// N is short name of Note func.
func (current *Node) N(note string) *Node {
	return current.Note(note)
}
