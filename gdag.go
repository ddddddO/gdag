package gdag

import (
	"fmt"
	"sort"
)

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
	// かつ、ソータブルな値が必須（ソートして使っているところがあるため）
	as    int
	note  string
	color string // done: #DarkGray

	gantt gantt

	upstream   []*Node
	downstream []*Node
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

	// TODO:
	// ここにクリティカルパス用の計算をいれるのは微妙な気がする
	tmpTerms := 0
	upstreamCritpath := path{}
	for _, up := range current.upstream {
		if up.gantt.criticalpath.terms > tmpTerms {
			tmpTerms = up.gantt.criticalpath.terms
			upstreamCritpath = up.gantt.criticalpath
		}
	}
	current.gantt.criticalpath = upstreamCritpath
	current.gantt.criticalpath.nodes = append(current.gantt.criticalpath.nodes, current)
	current.gantt.criticalpath.terms += current.gantt.term
	return current
}

func (current *Node) Note(note string) {
	current.note = note
}

func (current *Node) isDone() bool {
	return current.color == colorDone
}

func GenerateUML(node *Node) error {
	fmt.Println("@startuml")

	fmt.Println(generateComponents(node))
	fmt.Println(generateRelations(node))

	fmt.Println("@enduml")
	return nil
}

func generateComponents(node *Node) string {
	return generateComponent(node)
}

var uniqC = make(map[int]struct{})

func generateComponent(n *Node) string {
	if _, ok := uniqC[n.as]; ok {
		return ""
	}
	uniqC[n.as] = struct{}{}

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
		dst += generateComponent(d)
	}

	return dst
}

func generateRelations(node *Node) string {
	return generateRelation(node, "")
}

var uniqR = make(map[string]struct{})

func generateRelation(n *Node, out string) string {
	r := fmt.Sprintf("%d --> ", n.as)
	for _, d := range n.downstream {
		key := fmt.Sprintf("%d-%d", n.as, d.as)
		if _, ok := uniqR[key]; ok {
			continue
		}
		uniqR[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", r, d.as)
		out += generateRelation(d, tmp)
	}
	return out
}

// 要改修。本当はリレーションに基づいて生成した方が良さそうに思う。
func GenerateCheckList(n *Node) error {
	uniqAs(n)
	sorted := sortComponentList(uniqAS)

	out := ""
	for _, node := range sorted {
		if node.isDone() {
			out += fmt.Sprintf("- [x] %s\n", node.text)
		} else {
			out += fmt.Sprintf("- [ ] %s\n", node.text)
		}
	}
	fmt.Print(out)

	return nil
}

var uniqAS = make(map[int]*Node)

func uniqAs(n *Node) {
	if _, ok := uniqAS[n.as]; ok {
		return
	}
	uniqAS[n.as] = n

	for _, d := range n.downstream {
		uniqAs(d)
	}
}

func sortComponentList(uniq map[int]*Node) []*Node {
	keys := make([]int, 0, len(uniq))
	for k := range uniq {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sorted := make([]*Node, 0, len(keys))
	for _, k := range keys {
		v := uniq[k]
		sorted = append(sorted, v)
	}

	return sorted
}
