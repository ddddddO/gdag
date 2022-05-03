package gdag

import (
	"fmt"
)

// UML outputs dag PlantUML format.
func (start *Node) UML() (string, error) {
	ug := newUMLGenerator()

	ret := "@startuml" + "\n"
	ret += ug.generateComponents(start) + "\n"
	ret += ug.generateRelations(start) + "\n"
	ret += "@enduml"
	return ret, nil
}

type umlGenerator struct {
	uniqueC map[int]struct{}
	uniqueR map[string]struct{}
}

func newUMLGenerator() *umlGenerator {
	return &umlGenerator{
		uniqueC: map[int]struct{}{},
		uniqueR: map[string]struct{}{},
	}
}

func (ug *umlGenerator) generateComponents(start *Node) string {
	return ug.generateComponent(start)
}

func (ug *umlGenerator) generateComponent(n *Node) string {
	if _, ok := ug.uniqueC[n.index]; ok {
		return ""
	}
	ug.uniqueC[n.index] = struct{}{}

	ret := ""
	switch n.nodeType {
	case rectangle, usecase:
		s := fmt.Sprintf("%s \"%s\" as %d", n.nodeType, n.text, n.index)
		if len(n.color) != 0 {
			s += fmt.Sprintf(" %s", n.color)
		}
		s += "\n"
		ret += s
	}
	if len(n.note) != 0 {
		ret += fmt.Sprintf("note left\n%s\nend note\n", n.note)
	}

	for _, d := range n.downstream {
		ret += ug.generateComponent(d)
	}
	return ret
}

func (ug *umlGenerator) generateRelations(start *Node) string {
	return ug.generateRelation(start, "")
}

func (ug *umlGenerator) generateRelation(n *Node, out string) string {
	r := fmt.Sprintf("%d --> ", n.index)
	ret := out
	for _, d := range n.downstream {
		key := fmt.Sprintf("%d-%d", n.index, d.index)
		if _, ok := ug.uniqueR[key]; ok {
			continue
		}
		ug.uniqueR[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", r, d.index)
		ret += ug.generateRelation(d, tmp)
	}
	return ret
}
