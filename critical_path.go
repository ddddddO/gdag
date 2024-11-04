package gdag

type criticalPath struct {
	path    map[int]struct{} // key は Node の index
	sumHour float64
}

func (cp *criticalPath) contains(n *Node) bool {
	_, ok := cp.path[n.index]
	return ok
}

type criticalPathCalculator struct {
	allPaths [][]*Node // start からすべてのパス
}

func newCriticalPathCalculator() *criticalPathCalculator {
	return &criticalPathCalculator{}
}

func (cc *criticalPathCalculator) getCriticalPaths(start *Node) []*criticalPath {
	cc.walk(start, []*Node{})

	criticalPaths := []*criticalPath{}
	for _, path := range cc.allPaths {
		critical := &criticalPath{path: map[int]struct{}{}}
		for _, n := range path {
			critical.path[n.index] = struct{}{}
			critical.sumHour += n.hour
		}

		if critical.sumHour == 0 {
			continue
		}
		if len(criticalPaths) == 0 {
			criticalPaths = append(criticalPaths, critical)
			continue
		}
		if critical.sumHour == criticalPaths[0].sumHour {
			criticalPaths = append(criticalPaths, critical)
			continue
		}
		if critical.sumHour > criticalPaths[0].sumHour {
			criticalPaths = []*criticalPath{critical}
			continue
		}
	}

	return criticalPaths
}

func (cc *criticalPathCalculator) walk(current *Node, path []*Node) {
	path = append(path, current)

	if len(current.downstream) == 0 {
		cc.allPaths = append(cc.allPaths, path)
		return
	}

	for _, n := range current.downstream {
		cc.walk(n, path)
	}
}
