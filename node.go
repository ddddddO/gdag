package gdag

type Node struct {
	nodeType     nodeType
	index        int // mermaidの識別子としても利用する
	text         string
	note         string
	color        string // done: #DarkGray
	colorMermaid string // done: doneColor

	parent     *Node // TODO: 現状、中間ノードのためにおいてる
	downstream []*Node
}

type nodeType string

const (
	intermediate = nodeType("intermediate node")
	rectangle    = nodeType("rectangle")
	usecase      = nodeType("usecase")
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
		index:    nodeIdx,
	}
}

const (
	colorDone        = "#DarkGray"
	colorDoneMermaid = "doneColor"
)

func Done(nodes ...*Node) {
	for _, n := range nodes {
		n.color = colorDone
		n.colorMermaid = colorDoneMermaid
	}
}

func (upstream *Node) Con(current *Node) *Node {
	if upstream.nodeType == intermediate {
		for i := range upstream.downstream {
			upstream.downstream[i].downstream = append(upstream.downstream[i].downstream, current)
		}
		return current
	}

	for _, d := range upstream.downstream {
		if current.index == d.index {
			return d
		}
	}

	upstream.downstream = append(upstream.downstream, current)
	return current
}

func (upstream *Node) Fanout(nodes ...*Node) *Node {
	if upstream.nodeType == intermediate {
		// TODO: 後ほど、Fanout(n1, n2).Fanout(n3, n4)みたいに、中間ノードでもチェイン出来るようにしたい
		panic("Fanin/Fanout calls are not supported for intermediate node.")
	}

	intermediateNode := newNode(intermediate, "not output")
	intermediateNode.downstream = nodes

	intermediateNode.parent = upstream
	upstream.downstream = append(upstream.downstream, intermediateNode)
	return intermediateNode
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
