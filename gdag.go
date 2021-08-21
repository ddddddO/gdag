package gdag

import "fmt"

type nodeType string

const (
	rectangle = nodeType("rectangle")
	usecase   = nodeType("usecase")
)

type Node struct {
	nodeType nodeType
	text     string

	// plantumlでの識別子。自動で生成したいけど、例えばa, b, ...とかにしちゃうと、他の人がplantumlを編集するとき辛くなる。あと、識別子なので重複する場合はエラーとする。
	// の予定だったが、自動で生成する。連番。たぶん他の人がumlいじることはないとも思う。
	// せめて、連番ではなく、置換しやすいように少し長めのユニークなIDにしたほうがいいかも。テストできそうかも考えて実装した方が良さそう。
	as       int
	note     string
	color    string // done: #DarkGray
	parents  []*Node
	children []*Node
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

func Goal(text string) *Node {
	return newNode(rectangle, text)
}

func Task(text string) *Node {
	return newNode(usecase, text)
}

func Done(nodes ...*Node) {
	for _, n := range nodes {
		n.color = "#DarkGray"
	}
}

func (upstream *Node) Con(current *Node) *Node {
	for _, child := range upstream.children {
		if current.as == child.as {
			return child
		}
	}

	upstream.children = append(upstream.children, current)
	current.parents = append(current.parents, upstream)
	return current
}

func (current *Node) Note(note string) {
	current.note = note
}

func GenerateUML(node *Node) error {
	fmt.Println("@startuml")

	fmt.Println(generateComponents(node))
	fmt.Println(generateRelations(node))

	fmt.Println("@enduml")
	return nil
}

func generateComponents(node *Node) string {
	generateComponent(node)
	return dstComponents
}

var dstComponents string // 要リファクタ
func generateComponent(n *Node) {
	if n.nodeType == rectangle {
		s := fmt.Sprintf("rectangle \"%s\" as %d", n.text, n.as)
		if len(n.color) != 0 {
			s += fmt.Sprintf(" %s", n.color)
		}
		s += "\n"

		dstComponents += s
	}
	if n.nodeType == usecase {
		s := fmt.Sprintf("usecase \"%s\" as %d", n.text, n.as)
		if len(n.color) != 0 {
			s += fmt.Sprintf(" %s", n.color)
		}
		s += "\n"

		dstComponents += s
	}
	if len(n.note) != 0 {
		dstComponents += fmt.Sprintf("note left\n%s\nend note\n", n.note)
	}

	for _, child := range n.children {
		generateComponent(child)
	}
}

var dstRelations string // 要リファクタ
func generateRelations(node *Node) string {
	generateRelation(node)
	return dstRelations
}

var uniq = make(map[string]struct{})

func generateRelation(n *Node) {
	r := fmt.Sprintf("%d --> ", n.as)
	for _, child := range n.children {
		key := fmt.Sprintf("%d%d", n.as, child.as)
		if _, ok := uniq[key]; ok {
			continue
		}
		uniq[key] = struct{}{}

		dstRelations += fmt.Sprintf("%s%d\n", r, child.as)
		generateRelation(child)
	}
}
