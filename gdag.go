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
	as       string // plantumlでの識別子。自動で生成したいけど、例えばa, b, ...とかにしちゃうと、他の人がplantumlを編集するとき辛くなる。あと、識別子なので重複する場合はエラーとする。
	note     string
	parents  []*Node
	children []*Node
}

func newNode(nodeType nodeType, text, as string) *Node {
	return &Node{
		nodeType: nodeType,
		text:     text,
		as:       as,
	}
}

func NewGoal(text, as string) *Node {
	return newNode(rectangle, text, as)
}

func NewTask(text, as string) *Node {
	return newNode(usecase, text, as)
}

func (upstream *Node) Connect(current *Node) *Node {
	// ここはもう少し考える必要がありそう。
	for _, child := range upstream.children {
		if current.as == child.as {
			//panic("duplicate error")
			return current
		}
	}

	upstream.children = append(upstream.children, current)
	current.parents = append(current.parents, upstream)
	return current
}

func (current *Node) AddNote(note string) {
	current.note = note
}

func GenerateUML(goal *Node) error {
	fmt.Println("@startuml")

	fmt.Println(generateComponents(goal))
	fmt.Println(generateRelations(goal))

	fmt.Println("@enduml")
	return nil
}

func generateComponents(goal *Node) string {
	generateComponent(goal)
	return dstComponents
}

var dstComponents string // 要リファクタ
func generateComponent(n *Node) {
	if n.nodeType == rectangle {
		dstComponents += fmt.Sprintf("rectangle \"%s\" as %s\n", n.text, n.as)
	}
	if n.nodeType == usecase {
		dstComponents += fmt.Sprintf("usecase \"%s\" as %s\n", n.text, n.as)
	}
	if len(n.note) != 0 {
		dstComponents += fmt.Sprintf("note left\n%s\nend note\n", n.note)
	}

	for _, child := range n.children {
		generateComponent(child)
	}
}

var dstRelations string // 要リファクタ
func generateRelations(goal *Node) string {
	generateRelation(goal)
	return dstRelations
}

var uniq = make(map[string]struct{})

func generateRelation(n *Node) {
	r := fmt.Sprintf("%s --> ", n.as)
	for _, child := range n.children {
		key := n.as + child.as
		if _, ok := uniq[key]; ok {
			continue
		}
		uniq[key] = struct{}{}

		dstRelations += r + child.as + "\n"
		generateRelation(child)
	}
}
