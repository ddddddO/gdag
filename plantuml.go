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

func (ug *umlGenerator) generateComponent(node *Node) string {
	if _, ok := ug.uniqueC[node.index]; ok {
		return ""
	}
	ug.uniqueC[node.index] = struct{}{}

	ret := ""
	switch node.nodeType {
	case rectangle, usecase:
		s := fmt.Sprintf("%s \"%s\" as %d", node.nodeType, node.text, node.index)
		if len(node.color) != 0 {
			s += fmt.Sprintf(" %s", node.color)
		}
		s += "\n"
		ret += s
	}
	if len(node.note) != 0 {
		ret += fmt.Sprintf("note left\n%s\nend note\n", node.note)
	}

	for _, d := range node.downstream {
		ret += ug.generateComponent(d)
	}
	return ret
}

func (ug *umlGenerator) generateRelations(start *Node) string {
	return ug.generateRelation(start, "")
}

func (ug *umlGenerator) generateRelation(node *Node, out string) string {
	edge := fmt.Sprintf("%d --> ", node.index)
	ret := out
	for _, d := range node.downstream {
		key := fmt.Sprintf("%d-%d", node.index, d.index)
		if _, ok := ug.uniqueR[key]; ok {
			continue
		}
		ug.uniqueR[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", edge, d.index)
		ret += ug.generateRelation(d, tmp)
	}
	return ret
}
