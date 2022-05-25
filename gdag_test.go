package gdag_test

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func Example() {
	var dag *g.Node = g.DAG("ゴール(目的)")

	var design *g.Node = g.Task("設計")
	reviewDesign := g.Task("レビュー対応")

	developFeature1 := g.Task("feature1開発")
	developFeature1.Note("xxが担当")
	reviewDevelopFeature1 := g.Task("レビュー対応")

	developFeature2 := g.Task("feature2開発").Note("yyが担当")
	reviewDevelopFeature2 := g.Task("レビュー対応")

	prepareInfra := g.Task("インフラ準備").Note("zzが担当")

	test := g.Task("結合テスト")
	release := g.Task("リリース")
	finish := g.Task("finish")

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	uml, err := dag.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as 23
	// usecase "設計" as 24 #DarkGray
	// usecase "レビュー対応" as 25 #DarkGray
	// usecase "feature1開発" as 26
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 27
	// usecase "結合テスト" as 31
	// usecase "リリース" as 32
	// usecase "finish" as 33 #DarkGray
	// usecase "feature2開発" as 28 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 29
	// usecase "インフラ準備" as 30
	// note left
	// zzが担当
	// end note
	//
	// 23 --> 24
	// 24 --> 25
	// 25 --> 26
	// 26 --> 27
	// 27 --> 31
	// 31 --> 32
	// 32 --> 33
	// 25 --> 28
	// 28 --> 29
	// 29 --> 31
	// 25 --> 30
	// 30 --> 31
	//
	// @enduml
}
