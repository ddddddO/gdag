package gdag_test

import (
	"fmt"
	"testing"

	g "github.com/ddddddO/gdag"
)

func TestFaninFanout_Continuously(t *testing.T) {
	dag := g.DAG("Fanin/Fanout")
	dag.Fanout(
		g.Task("t1"), g.Task("t2"),
	).Fanin(
		g.Task("t3"),
	).Fanout(
		g.Task("t4"), g.Task("t5"), g.Task("t6"), g.Task("t7"),
	).Fanin(
		g.Task("t8"),
	).Con(
		g.Task("t9"),
	).Fanout(
		g.Task("t10"), g.Task("t11"),
	).Fanin(
		g.Task("end"),
	)
	uml, err := dag.UML()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uml)
}
