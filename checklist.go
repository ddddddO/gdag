package gdag

import (
	"fmt"
	"sort"
)

// CheckList outputs task check list.
func (start *Node) CheckList() (string, error) {
	clg := newCheckListGenerator()

	clg.makeUnique(start)
	sorted := clg.sortComponentList()

	ret := ""
	for _, node := range sorted {
		if node.isDAG() {
			ret += fmt.Sprintf("### %s\n", node.text)
			continue
		}

		if node.isDone() {
			ret += fmt.Sprintf("- [x] %s\n", node.text)
		} else {
			ret += fmt.Sprintf("- [ ] %s\n", node.text)
		}
	}
	return ret, nil
}

type checkListGenerator struct {
	unique map[int]*Node
}

func newCheckListGenerator() *checkListGenerator {
	return &checkListGenerator{
		unique: map[int]*Node{},
	}
}

func (clg *checkListGenerator) makeUnique(node *Node) {
	if _, ok := clg.unique[node.index]; ok {
		return
	}
	clg.unique[node.index] = node

	for _, d := range node.downstream {
		clg.makeUnique(d)
	}
}

func (clg *checkListGenerator) sortComponentList() []*Node {
	keys := make([]int, 0, len(clg.unique))
	for k := range clg.unique {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	sorted := make([]*Node, 0, len(keys))
	for _, k := range keys {
		v := clg.unique[k]
		sorted = append(sorted, v)
	}
	return sorted
}
