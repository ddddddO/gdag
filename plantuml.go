package gdag

import (
	"fmt"
)

// UML outputs dag PlantUML format.
func (start *Node) UML() (string, error) {
	start.startPoint = true
	cc := newCriticalPathCalculator()
	ug := newUMLGenerator(cc.getCriticalPaths(start))

	ret := "@startuml" + "\n"
	ret += ug.generateComponents(start) + "\n"
	ret += ug.generateRelations(start) + "\n"
	ret += "@enduml"

	start.startPoint = false
	return ret, nil
}

// UMLNoCritical outputs dag PlantUML format that does not represent critical path.
func (start *Node) UMLNoCritical() (string, error) {
	start.startPoint = true
	ug := newUMLGenerator(nil)

	ret := "@startuml" + "\n"
	ret += ug.generateComponents(start) + "\n"
	ret += ug.generateRelations(start) + "\n"
	ret += "@enduml"

	start.startPoint = false
	return ret, nil
}

type umlGenerator struct {
	criticalPaths    []*criticalPath
	uniqueComponents map[int]struct{}
	uniqueRelations  map[string]struct{}
}

func newUMLGenerator(criticalPaths []*criticalPath) *umlGenerator {
	return &umlGenerator{
		criticalPaths:    criticalPaths,
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

	ret := ug.renderComponent(node)
	for _, d := range node.downstream {
		ret += ug.generateComponent(d)
	}
	return ret
}

func (ug *umlGenerator) renderComponent(node *Node) string {
	ret := ""
	switch node.nodeType {
	case rectangle, usecase:
		s := fmt.Sprintf("%s \"%s\" as %d", node.nodeType, node.text, node.index)
		if node.hour > 0 {
			s = fmt.Sprintf("%s \"%s (%.1fh)\" as %d", node.nodeType, node.text, node.hour, node.index)
		}

		if len(node.color) != 0 {
			s += fmt.Sprintf(" %s", node.color)
			if ug.isCritical(node) && !node.isDAG() {
				s += fmt.Sprintf("-%s", "Yellow")
			}
		} else {
			if ug.isCritical(node) && !node.isDAG() {
				s += fmt.Sprintf(" %s", "#Yellow")
			}
		}

		s += "\n"
		ret += s
	}
	if len(node.note) != 0 {
		ret += fmt.Sprintf("note left\n%s\nend note\n", node.note)
	}
	return ret
}

func (ug *umlGenerator) isCritical(current *Node) bool {
	for _, cp := range ug.criticalPaths {
		if cp.contains(current) {
			return true
		}
	}
	return false
}

func (ug *umlGenerator) generateRelations(start *Node) string {
	return ug.generateRelation(start, "")
}

func (ug *umlGenerator) generateRelation(node *Node, out string) string {
	edge := fmt.Sprintf("%d --> ", node.index)
	ret := out
	for _, d := range node.downstream {
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
