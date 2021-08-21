package gdag_test

import (
	g "github.com/ddddddO/gdag"
)

func Example() {
	goal := g.Goal("ゴール(目的)")

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

	goal.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	if err := g.GenerateUML(goal); err != nil {
		panic(err)
	}
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as 1
	// usecase "設計" as 2 #DarkGray
	// usecase "レビュー対応" as 3 #DarkGray
	// usecase "feature1開発" as 4
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 5
	// usecase "結合テスト" as 9
	// usecase "リリース" as 10
	// usecase "finish" as 11 #DarkGray
	// usecase "feature2開発" as 6 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 7
	// usecase "インフラ準備" as 8
	// note left
	// zzが担当
	// end note
	//
	// 1 --> 2
	// 2 --> 3
	// 3 --> 4
	// 4 --> 5
	// 5 --> 9
	// 9 --> 10
	// 10 --> 11
	// 3 --> 6
	// 6 --> 7
	// 7 --> 9
	// 3 --> 8
	// 8 --> 9
	//
	// @enduml
}

func ExampleGenerateCheckList() {
	goal := g.Goal("ゴール(目的)")

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

	goal.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	if err := g.GenerateCheckList(design); err != nil {
		panic(err)
	}
	// Output:
	// - [ ] 設計
	// - [ ] レビュー対応
	// - [ ] feature1開発
	// - [ ] レビュー対応
	// - [ ] feature2開発
	// - [ ] レビュー対応
	// - [ ] インフラ準備
	// - [ ] 結合テスト
	// - [ ] リリース
	// - [ ] finish
}
