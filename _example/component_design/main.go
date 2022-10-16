package main

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func main() {
	mainComp := g.DAG("Main")
	view := g.Task("View")
	controllers := g.Task("Controllers")
	presenters := g.Task("Presenters")
	interactors := g.Task("Interactors")
	authorizer := g.Task("Authorizer")
	database := g.Task("Database")
	entities := g.Task("Entities")
	permissions := g.Task("Permissions")

	mainComp.Con(view).Con(presenters).Con(interactors)
	mainComp.Con(interactors).Con(entities)
	mainComp.Con(database).Con(interactors)
	database.Con(entities).Con(permissions)
	mainComp.Con(controllers).Con(interactors).Con(entities)
	mainComp.Con(authorizer).Con(interactors)
	authorizer.Con(permissions)

	uml, err := mainComp.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
}