package gdag

import (
	"fmt"
)

// Mermaid outputs dag mermaidjs.
func (start *Node) Mermaid() (string, error) {
	mg := newMermaidGenerator()

	ret := "graph TD" + "\n"
	ret += "classDef doneColor fill:#868787" + "\n"
	ret += mg.generateComponents(start) + "\n"
	ret += mg.generateRelations(start) + "\n"
	return ret, nil
}

type mermaidGenerator struct {
	uniqueC map[int]struct{}
	uniqueR map[string]struct{}
}

func newMermaidGenerator() *mermaidGenerator {
	return &mermaidGenerator{
		uniqueC: map[int]struct{}{},
		uniqueR: map[string]struct{}{},
	}
}

func (mg *mermaidGenerator) generateComponents(start *Node) string {
	return mg.generateComponent(start)
}

func (mg *mermaidGenerator) generateComponent(n *Node) string {
	if _, ok := mg.uniqueC[n.index]; ok {
		return ""
	}
	mg.uniqueC[n.index] = struct{}{}

	ret := ""
	// TODO: mermaid用に修正するかどうか。リファクタは必要
	switch n.nodeType {
	case rectangle:
		s := fmt.Sprintf("%d(\"%s\")", n.index, n.text)
		if len(n.colorMermaid) != 0 {
			s += fmt.Sprintf(":::%s", n.colorMermaid)
		}
		s += "\n"
		ret += s
	case usecase:
		s := fmt.Sprintf("%d([\"%s\"])", n.index, n.text)
		if len(n.colorMermaid) != 0 {
			s += fmt.Sprintf(":::%s", n.colorMermaid)
		}
		s += "\n"
		ret += s
	}
	if len(n.note) != 0 {
		// noop
	}

	for _, d := range n.downstream {
		ret += mg.generateComponent(d)
	}
	return ret
}

func (mg *mermaidGenerator) generateRelations(start *Node) string {
	return mg.generateRelation(start, "")
}

func (mg *mermaidGenerator) generateRelation(n *Node, out string) string {
	r := fmt.Sprintf("%d --> ", n.index)
	ret := out
	for _, d := range n.downstream {
		key := fmt.Sprintf("%d-%d", n.index, d.index)
		if _, ok := mg.uniqueR[key]; ok {
			continue
		}
		mg.uniqueR[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", r, d.index)
		ret += mg.generateRelation(d, tmp)
	}
	return ret
}