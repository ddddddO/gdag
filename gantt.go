package gdag

import (
	"fmt"
	"sort"
)

type gantt struct {
	title      string // text of Node
	identifier int    // as of Node
	action     string // e.g. active/done ...

	// startDateに空文字以外の日付が入っていれば、startDate + , + term + d
	// startDateが空文字であれば、 after + 上流タスクのidentifier + , + term + d
	startDate    string // e.g. 2014-01-07
	term         int    // e.g. 3(days) -> after upstream, 3d
	criticalpath path
}

type path struct {
	nodes []*Node
	terms int
}

func newGantt(title string, identifier int, startDate string, term int) gantt {
	return gantt{
		title:      title,
		identifier: identifier,
		startDate:  startDate,
		term:       term,
	}
}

// WithGanttStart is ゴール(rectangle)直下のタスク(usecase)に一回呼び出す。
// その後ろのタスクはWithGantt funcを呼び出す。なんか微妙な
func (n *Node) WithGanttStart(startDate string, term int) {
	n.gantt = newGantt(n.text, n.as, startDate, term)
}

func (n *Node) WithGantt(term int) {
	n.gantt = newGantt(n.text, n.as, "", term)
}

func GenerateGantt(node *Node) error {
	gantt := generateGantt(node)
	out := fmt.Sprintf(ganttTemplate, gantt)
	fmt.Println(out)
	return nil
}

func generateGantt(node *Node) string {
	uniqGantt(node)
	sorted := sortGantt(uniqG)
	return generateG(sorted)
}

var uniqG = make(map[int]*Node)

func uniqGantt(n *Node) {
	if _, ok := uniqG[n.as]; ok {
		return
	}
	uniqG[n.as] = n

	for _, d := range n.downstream {
		uniqGantt(d)
	}
}

func sortGantt(uniq map[int]*Node) []*Node {
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

func generateG(sorted []*Node) string {
	out := ""
	for _, n := range sorted {
		tmp := ""
		if n.nodeType == rectangle {
			tmp += "section " + n.text
		}
		if n.nodeType == usecase {
			tmp += "\t" + n.gantt.title + " :" + fmt.Sprintf("%d", n.gantt.identifier) + ","
			// TODO: このあたりにactionいれる
			if len(n.gantt.startDate) != 0 {
				tmp += n.gantt.startDate + ","
			} else {
				preNodeIdx := len(n.gantt.criticalpath.nodes) - 2
				after := n.gantt.criticalpath.nodes[preNodeIdx].gantt
				tmp += fmt.Sprintf("after %d", after.identifier) + ","
			}
			tmp += fmt.Sprintf("%dd", n.gantt.term)
		}
		tmp += "\n"
		out += tmp
	}
	return out
}
