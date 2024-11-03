package gdag

type Node struct {
	nodeType nodeType
	index    int // mermaidの識別子としても利用する
	text     string
	note     string
	hour     float64 // 見積時間

	color        string // done: #DarkGray
	colorMermaid string // done: doneColor

	// parent     *Node // TODO: 現状、中間ノードのためにおいてる
	downstream []*Node
	startPoint bool // レンダリング処理の最初の node ということ
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
	for _, d := range upstream.downstream {
		if current.index == d.index {
			return d
		}
	}

	upstream.downstream = append(upstream.downstream, current)
	return current
}

type FanIO struct {
	upstream *Node
}

func (upstream *Node) Fanout(nodes ...*Node) *FanIO {
	upstream.downstream = append(upstream.downstream, nodes...)
	return &FanIO{
		upstream: upstream,
	}
}

func (fio *FanIO) Fanin(current *Node) *Node {
	for i := range fio.upstream.downstream {
		fio.upstream.downstream[i].downstream = append(fio.upstream.downstream[i].downstream, current)
	}
	return current
}

func (current *Node) Note(note string) *Node {
	current.note = note
	return current
}

func (current *Node) Hour(hour float64) *Node {
	current.hour = hour
	return current
}

func (current *Node) isDone() bool {
	return current.color == colorDone
}

func (current *Node) isDAG() bool {
	return current.nodeType == rectangle
}
