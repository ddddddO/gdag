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
	uniqueComponents map[int]struct{}
	uniqueRelations  map[string]struct{}
}

func newMermaidGenerator() *mermaidGenerator {
	return &mermaidGenerator{
		uniqueComponents: map[int]struct{}{},
		uniqueRelations:  map[string]struct{}{},
	}
}

func (mg *mermaidGenerator) generateComponents(start *Node) string {
	return mg.generateComponent(start)
}

func (mg *mermaidGenerator) generateComponent(node *Node) string {
	if _, ok := mg.uniqueComponents[node.index]; ok {
		return ""
	}
	mg.uniqueComponents[node.index] = struct{}{}

	ret := (*mermaidGenerator)(nil).renderComponent(node)
	for _, d := range node.downstream {
		ret += mg.generateComponent(d)
	}
	return ret
}

func (*mermaidGenerator) renderComponent(node *Node) string {
	ret := ""
	// TODO: mermaid用に修正するかどうか。リファクタは必要
	switch node.nodeType {
	case rectangle:
		s := fmt.Sprintf("%d(\"%s\")", node.index, node.text)
		if len(node.colorMermaid) != 0 {
			s += fmt.Sprintf(":::%s", node.colorMermaid)
		}
		s += "\n"
		ret += s
	case usecase:
		s := fmt.Sprintf("%d([\"%s\"])", node.index, node.text)
		if len(node.colorMermaid) != 0 {
			s += fmt.Sprintf(":::%s", node.colorMermaid)
		}
		s += "\n"
		ret += s
	}
	// nolint:staticcheck
	if len(node.note) != 0 {
		// noop
	}
	return ret
}

func (mg *mermaidGenerator) generateRelations(start *Node) string {
	return mg.generateRelation(start, "")
}

func (mg *mermaidGenerator) generateRelation(node *Node, out string) string {
	edge := fmt.Sprintf("%d --> ", node.index)
	ret := out
	for _, d := range node.downstream {
		key := fmt.Sprintf("%d-%d", node.index, d.index)
		if _, ok := mg.uniqueRelations[key]; ok {
			continue
		}
		mg.uniqueRelations[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", edge, d.index)
		ret += mg.generateRelation(d, tmp)
	}
	return ret
}
