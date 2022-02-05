package gdag

import (
	"fmt"
	"sort"
)

type Node struct {
	nodeType nodeType
	text     string
	// plantumlでの識別子。自動で生成したいけど、例えばa, b, ...とかにしちゃうと、他の人がplantumlを編集するとき辛くなる。あと、識別子なので重複する場合はエラーとする。
	// の予定だったが、自動で生成する。連番。たぶん他の人がumlいじることはないとも思う。
	// せめて、連番ではなく、置換しやすいように少し長めのユニークなIDにしたほうがいいかも。テストできそうかも考えて実装した方が良さそう。
	// かつ、ソータブルな値が必須（ソートして使っているところがあるため）
	as    int
	note  string
	color string // done: #DarkGray

	upstream   []*Node
	downstream []*Node
}

type nodeType string

const (
	rectangle = nodeType("rectangle")
	usecase   = nodeType("usecase")
)

func DAG(text string) *Node {
	return newNode(rectangle, text)
}

func Task(text string) *Node {
	return newNode(usecase, text)
}

var nodeIdx int

func newNode(nodeType nodeType, text string) *Node {
	nodeIdx++

	return &Node{
		nodeType: nodeType,
		text:     text,
		as:       nodeIdx,
	}
}

const (
	colorDone = "#DarkGray"
)

func Done(nodes ...*Node) {
	for _, n := range nodes {
		n.color = colorDone
	}
}

func (upstream *Node) Con(current *Node) *Node {
	for _, d := range upstream.downstream {
		if current.as == d.as {
			return d
		}
	}

	upstream.downstream = append(upstream.downstream, current)
	current.upstream = append(current.upstream, upstream)

	return current
}

func (current *Node) Note(note string) {
	current.note = note
}

func (current *Node) isDone() bool {
	return current.color == colorDone
}

func (current *Node) isDAG() bool {
	return current.nodeType == rectangle
}

// UML outputs dag plant UML.
func (start *Node) UML() (string, error) {
	ug := newUMLGenerator()

	ret := "@startuml" + "\n"
	ret += ug.generateComponents(start) + "\n"
	ret += ug.generateRelations(start) + "\n"
	ret += "@enduml"
	return ret, nil
}

type umlGenerator struct {
	uniqueC map[int]struct{}
	uniqueR map[string]struct{}
}

func newUMLGenerator() *umlGenerator {
	return &umlGenerator{
		uniqueC: map[int]struct{}{},
		uniqueR: map[string]struct{}{},
	}
}

func (ug *umlGenerator) generateComponents(start *Node) string {
	return ug.generateComponent(start)
}

func (ug *umlGenerator) generateComponent(n *Node) string {
	if _, ok := ug.uniqueC[n.as]; ok {
		return ""
	}
	ug.uniqueC[n.as] = struct{}{}

	dst := ""
	switch n.nodeType {
	case rectangle, usecase:
		s := fmt.Sprintf("%s \"%s\" as %d", n.nodeType, n.text, n.as)
		if len(n.color) != 0 {
			s += fmt.Sprintf(" %s", n.color)
		}
		s += "\n"

		dst += s
	}
	if len(n.note) != 0 {
		dst += fmt.Sprintf("note left\n%s\nend note\n", n.note)
	}

	for _, d := range n.downstream {
		dst += ug.generateComponent(d)
	}

	return dst
}

func (ug *umlGenerator) generateRelations(start *Node) string {
	return ug.generateRelation(start, "")
}

func (ug *umlGenerator) generateRelation(n *Node, out string) string {
	r := fmt.Sprintf("%d --> ", n.as)
	for _, d := range n.downstream {
		key := fmt.Sprintf("%d-%d", n.as, d.as)
		if _, ok := ug.uniqueR[key]; ok {
			continue
		}
		ug.uniqueR[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", r, d.as)
		out += ug.generateRelation(d, tmp)
	}
	return out
}

// CheckList outputs task check list.
func (start *Node) CheckList() (string, error) {
	clg := newCheckListGenerator()

	clg.makeUnique(start)
	sorted := clg.sortComponentList()

	ret := ""
	for _, node := range sorted {
		if node.isDAG() {
			ret += fmt.Sprintf("### %s\n", node.text)
			continue
		}

		if node.isDone() {
			ret += fmt.Sprintf("- [x] %s\n", node.text)
		} else {
			ret += fmt.Sprintf("- [ ] %s\n", node.text)
		}
	}
	return ret, nil
}

type checkListGenerator struct {
	unique map[int]*Node
}

func newCheckListGenerator() *checkListGenerator {
	return &checkListGenerator{
		unique: map[int]*Node{},
	}
}

func (clg *checkListGenerator) makeUnique(n *Node) {
	if _, ok := clg.unique[n.as]; ok {
		return
	}
	clg.unique[n.as] = n

	for _, d := range n.downstream {
		clg.makeUnique(d)
	}
}

func (clg *checkListGenerator) sortComponentList() []*Node {
	keys := make([]int, 0, len(clg.unique))
	for k := range clg.unique {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sorted := make([]*Node, 0, len(keys))
	for _, k := range keys {
		v := clg.unique[k]
		sorted = append(sorted, v)
	}

	return sorted
}
