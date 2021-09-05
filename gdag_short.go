package gdag

// G is short name of Goal func.
func G(text string) *Node {
	return Goal(text)
}

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
func (current *Node) N(note string) {
	current.Note(note)
}

// GUML is short name of GenerateUML func.
func GUML(node *Node) error {
	return GenerateUML(node)
}

// GCL is short name of GenerateCheckList func.
func GCL(node *Node) error {
	return GenerateCheckList(node)
}