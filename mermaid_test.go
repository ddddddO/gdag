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
	// Output:
	// graph TD
	// classDef doneColor fill:#868787
	// 34("ゴール(目的)")
	// 35(["設計"]):::doneColor
	// 36(["レビュー対応"]):::doneColor
	// 37(["feature1開発"])
	// 38(["レビュー対応"])
	// 42(["結合テスト"])
	// 43(["リリース"])
	// 44(["finish"]):::doneColor
	// 39(["feature2開発"]):::doneColor
	// 40(["レビュー対応"])
	// 41(["インフラ準備"])
	//
	// 34 --> 35
	// 35 --> 36
	// 36 --> 37
	// 37 --> 38
	// 38 --> 42
	// 42 --> 43
	// 43 --> 44
	// 36 --> 39
	// 39 --> 40
	// 40 --> 42
	// 36 --> 41
	// 41 --> 42

}
