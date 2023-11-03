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
	// Unordered output:
	// @startuml
	// rectangle "Fanin/Fanout" as 1
	// usecase "t1" as 2
	// usecase "t3" as 4
	// usecase "t4" as 5
	// usecase "t8" as 9
	// usecase "t9" as 10
	// usecase "t10" as 11
	// usecase "end" as 13
	// usecase "t11" as 12
	// usecase "t5" as 6
	// usecase "t6" as 7
	// usecase "t7" as 8
	// usecase "t2" as 3
	//
	// 1 --> 2
	// 2 --> 4
	// 4 --> 5
	// 5 --> 9
	// 9 --> 10
	// 10 --> 11
	// 11 --> 13
	// 10 --> 12
	// 12 --> 13
	// 4 --> 6
	// 6 --> 9
	// 4 --> 7
	// 7 --> 9
	// 4 --> 8
	// 8 --> 9
	// 1 --> 3
	// 3 --> 4
	//
	// @enduml
}
