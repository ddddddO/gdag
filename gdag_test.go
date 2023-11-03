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
	// Unordered output:
	// @startuml
	// rectangle "ゴール(目的)" as 36
	// usecase "設計" as 37 #DarkGray
	// usecase "レビュー対応" as 38 #DarkGray
	// usecase "feature1開発" as 39
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 40
	// usecase "結合テスト" as 44
	// usecase "リリース" as 45
	// usecase "finish" as 46 #DarkGray
	// usecase "feature2開発" as 41 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 42
	// usecase "インフラ準備" as 43
	// note left
	// zzが担当
	// end note
	//
	// 36 --> 37
	// 37 --> 38
	// 38 --> 39
	// 39 --> 40
	// 40 --> 44
	// 44 --> 45
	// 45 --> 46
	// 38 --> 41
	// 41 --> 42
	// 42 --> 44
	// 38 --> 43
	// 43 --> 44
	//
	// @enduml
}
