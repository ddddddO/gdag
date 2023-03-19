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
	uniqueComponents map[int]struct{}
	uniqueRelations  map[string]struct{}
}

func newUMLGenerator() *umlGenerator {
	return &umlGenerator{
		uniqueComponents: map[int]struct{}{},
		uniqueRelations:  map[string]struct{}{},
	}
}

func (ug *umlGenerator) generateComponents(start *Node) string {
	return ug.generateComponent(start)
}

func (ug *umlGenerator) generateComponent(node *Node) string {
	if _, ok := ug.uniqueComponents[node.index]; ok {
		return ""
	}
	ug.uniqueComponents[node.index] = struct{}{}

	ret := (*umlGenerator)(nil).renderComponent(node)
	for _, d := range node.downstream {
		ret += ug.generateComponent(d)
	}
	return ret
}

func (*umlGenerator) renderComponent(node *Node) string {
	ret := ""
	switch node.nodeType {
	case intermediate:
		break
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
	return ret
}

func (ug *umlGenerator) generateRelations(start *Node) string {
	return ug.generateRelation(start, "")
}

func (ug *umlGenerator) generateRelation(node *Node, out string) string {
	edge := fmt.Sprintf("%d --> ", node.index)
	if node.nodeType == intermediate {
		edge = fmt.Sprintf("%d --> ", node.parent.index)
		// parent := node.parent
		// for ; parent != nil && parent.nodeType == intermediate; {
		// 	parent = parent.parent
		// }
		// ret := out
		// for _, dd := range parent.downstream {
		// 	edge = fmt.Sprintf("%d --> ", dd.index)
		// 	// ret := out
		// 	for _, d := range node.downstream {
		// 		if d.nodeType == intermediate {
		// 			ret += ug.generateRelation(d, "")
		// 			continue
		// 		}

		// 		key := fmt.Sprintf("%d-%d", dd.index, d.index)
		// 		if _, ok := ug.uniqueRelations[key]; ok {
		// 			continue
		// 		}
		// 		ug.uniqueRelations[key] = struct{}{}

		// 		tmp := fmt.Sprintf("%s%d\n", edge, d.index)
		// 		ret += ug.generateRelation(d, tmp)
		// 	}
		// 	// return ret
		// }
		// return ret
	}

	ret := out
	for _, d := range node.downstream {
		if d.nodeType == intermediate {
			ret += ug.generateRelation(d, "")
			continue
		}

		key := fmt.Sprintf("%d-%d", node.index, d.index)
		if _, ok := ug.uniqueRelations[key]; ok {
			continue
		}
		ug.uniqueRelations[key] = struct{}{}

		tmp := fmt.Sprintf("%s%d\n", edge, d.index)
		ret += ug.generateRelation(d, tmp)
	}
	return ret
}
