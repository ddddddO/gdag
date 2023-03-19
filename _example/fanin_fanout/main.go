package main

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func main() {
	dag := g.DAG("Fanin/Fanout")
	task1 := g.Task("task1")
	task2 := g.Task("task2")
	task3 := g.Task("task3")
	task4 := g.Task("task4")

	// block1
	dag.Fanout(task1, task2, task3).Con(task4)
	task5 := task4.Con(g.Task("task5"))
	task6 := task4.Con(g.Task("task6"))
	task7 := task4.Con(g.Task("task7"))
	g.Fanin(task5, task6, task7).Con(g.Task("end"))

	// block2
	// dag.Fanout(task1, task2, task3).Con(task4).Fanout(g.Task("task5"), g.Task("task6"), g.Task("task7")).Con(g.Task("end"))

	uml, err := dag.Mermaid()
	// uml, err := dag.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
}