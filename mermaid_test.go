package gdag_test

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func ExampleNode_Mermaid() {
	var dag *g.Node = g.DAG("ゴール(目的)")

	var design *g.Node = g.Task("設計")
	reviewDesign := g.Task("レビュー対応")

	developFeature1 := g.Task("feature1開発")
	developFeature1.Note("noop")
	reviewDevelopFeature1 := g.Task("レビュー対応")

	developFeature2 := g.Task("feature2開発")
	developFeature2.Note("noop")
	reviewDevelopFeature2 := g.Task("レビュー対応")

	prepareInfra := g.Task("インフラ準備")
	prepareInfra.Note("noop")

	test := g.Task("結合テスト")
	release := g.Task("リリース")
	finish := g.Task("finish")

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	mermaid, err := dag.Mermaid()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(mermaid)
	// Unordered output:
	// graph TD
	// classDef doneColor fill:#868787
	// 47("ゴール(目的)")
	// 48(["設計"]):::doneColor
	// 49(["レビュー対応"]):::doneColor
	// 50(["feature1開発"])
	// 51(["レビュー対応"])
	// 55(["結合テスト"])
	// 56(["リリース"])
	// 57(["finish"]):::doneColor
	// 52(["feature2開発"]):::doneColor
	// 53(["レビュー対応"])
	// 54(["インフラ準備"])
	//
	// 47 --> 48
	// 48 --> 49
	// 49 --> 50
	// 50 --> 51
	// 51 --> 55
	// 55 --> 56
	// 56 --> 57
	// 49 --> 52
	// 52 --> 53
	// 53 --> 55
	// 49 --> 54
	// 54 --> 55
}
