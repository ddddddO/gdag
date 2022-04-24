package gdag

type Node struct {
	nodeType nodeType
	text     string
	// plantumlでの識別子。自動で生成したいけど、例えばa, b, ...とかにしちゃうと、他の人がplantumlを編集するとき辛くなる。あと、識別子なので重複する場合はエラーとする。
	// の予定だったが、自動で生成する。連番。たぶん他の人がumlいじることはないとも思う。
	// せめて、連番ではなく、置換しやすいように少し長めのユニークなIDにしたほうがいいかも。テストできそうかも考えて実装した方が良さそう。
	// かつ、ソータブルな値が必須（ソートして使っているところがあるため）
	as           int // mermaidjsの識別子としても利用する
	note         string
	color        string // done: #DarkGray
	colorMermaid string // done: doneColor

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
		n.colorMermaid = colorDoneMermaid
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

func (current *Node) Note(note string) *Node {
	current.note = note
	return current
}

func (current *Node) isDone() bool {
	return current.color == colorDone
}

func (current *Node) isDAG() bool {
	return current.nodeType == rectangle
}
