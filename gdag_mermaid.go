package gdag

import (
	"fmt"
)

const (
	colorDoneMermaid = "doneColor"
)

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

// Mermaid outputs dag mermaidjs.
func (start *Node) Mermaid() (string, error) {
	mg := newMermaidGenerator()

	ret := "graph TD" + "\n"
	ret += "classDef doneColor fill:#868787" + "\n"
	ret += mg.generateComponents(start) + "\n"
	ret += mg.generateRelations(start) + "\n"
	return ret, nil
}

func (mg *mermaidGenerator) generateComponents(start *Node) string {
	return mg.generateComponent(start)
}

func (mg *mermaidGenerator) generateComponent(n *Node) string {
	if _, ok := mg.uniqueC[n.as]; ok {
		return ""
	}
	mg.uniqueC[n.as] = struct{}{}

	dst := ""
	// TODO: mermaidjs用に修正するかどうか。リファクタは必要
	switch n.nodeType {
	case rectangle:
		s := fmt.Sprintf("%d(\"%s\")", n.as, n.text)
		if len(n.colorMermaid) != 0 {
			s += fmt.Sprintf(":::%s", n.colorMermaid)
		}
		s += "\n"

		dst += s
	case usecase:
		s := fmt.Sprintf("%d([\"%s\"])", n.as, n.text)
		if len(n.colorMermaid) != 0 {
			s += fmt.Sprintf(":::%s", n.colorMermaid)
		}
		s += "\n"

		dst += s
	}
	if len(n.note) != 0 {
		// noop
	}

	for _, d := range n.downstream {
		dst += mg.generateComponent(d)
	}

	return dst
}

func (mg *mermaidGenerator) generateRelations(start *Node) string {
	return mg.generateRelation(start, "")
}

func (mg *mermaidGenerator) generateRelation(n *Node, out string) string {
	r := fmt.Sprintf("%d --> ", n.as)
	for _, d := range n.downstream {
		key := fmt.Sprintf("%d-%d", n.as, d.as)
		if _, ok := mg.uniqueR[key]; ok {
			continue
		}
		mg.uniqueR[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", r, d.as)
		out += mg.generateRelation(d, tmp)
	}
	return out
}
