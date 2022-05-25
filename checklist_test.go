package gdag_test

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func ExampleNode_CheckList() {
	dag := g.DAG("ゴール(目的)")

	design := g.Task("設計")
	reviewDesign := g.Task("レビュー対応")

	developFeature1 := g.Task("feature1開発")
	developFeature1.Note("xxが担当")
	reviewDevelopFeature1 := g.Task("レビュー対応")

	developFeature2 := g.Task("feature2開発")
	developFeature2.Note("yyが担当")
	reviewDevelopFeature2 := g.Task("レビュー対応")

	prepareInfra := g.Task("インフラ準備")
	prepareInfra.Note("zzが担当")

	test := g.Task("結合テスト")
	release := g.Task("リリース")
	finish := g.Task("finish")

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	checkList, err := dag.CheckList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(checkList)
	// Output:
	// ### ゴール(目的)
	// - [x] 設計
	// - [x] レビュー対応
	// - [ ] feature1開発
	// - [ ] レビュー対応
	// - [x] feature2開発
	// - [ ] レビュー対応
	// - [ ] インフラ準備
	// - [ ] 結合テスト
	// - [ ] リリース
	// - [x] finish
}

func ExampleNode_CheckList_Multiple() {
	dag := g.DAG("ゴール(目的)")

	design := g.Task("設計")
	reviewDesign := g.Task("レビュー対応")

	developFeature1 := g.Task("feature1開発")
	developFeature1.Note("xxが担当")
	reviewDevelopFeature1 := g.Task("レビュー対応")

	developFeature2 := g.Task("feature2開発")
	developFeature2.Note("yyが担当")
	reviewDevelopFeature2 := g.Task("レビュー対応")

	prepareInfra := g.Task("インフラ準備")
	prepareInfra.Note("zzが担当")

	test := g.Task("結合テスト")
	release := g.Task("リリース")
	finish := g.Task("finish")

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	dagCheckList, err := dag.CheckList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(dagCheckList)

	infraCheckList, err := prepareInfra.CheckList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(infraCheckList)
	// Output:
	// ### ゴール(目的)
	// - [x] 設計
	// - [x] レビュー対応
	// - [ ] feature1開発
	// - [ ] レビュー対応
	// - [x] feature2開発
	// - [ ] レビュー対応
	// - [ ] インフラ準備
	// - [ ] 結合テスト
	// - [ ] リリース
	// - [x] finish
	//
	// - [ ] インフラ準備
	// - [ ] 結合テスト
	// - [ ] リリース
	// - [x] finish
}
