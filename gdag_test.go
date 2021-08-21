package gdag_test

import (
	g "github.com/ddddddO/gdag"
)

func Example() {
	goal := g.Goal("ゴール(目的)")

	design := g.Task("設計")
	review_design := g.Task("レビュー対応")

	develop_feature_1 := g.Task("feature1開発")
	develop_feature_1.AddNote("xxが担当")
	review_develop_feature_1 := g.Task("レビュー対応")

	develop_feature_2 := g.Task("feature2開発")
	develop_feature_2.AddNote("yyが担当")
	review_develop_feature_2 := g.Task("レビュー対応")

	test := g.Task("結合テスト")
	release := g.Task("リリース")
	finish := g.Task("finish")

	goal.Con(design).Con(review_design).Con(develop_feature_1).Con(review_develop_feature_1).Con(test)
	review_design.Con(develop_feature_2).Con(review_develop_feature_2).Con(test)
	test.Con(release).Con(finish)

	if err := g.GenerateUML(goal); err != nil {
		panic(err)
	}
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as 1
	// usecase "設計" as 2
	// usecase "レビュー対応" as 3
	// usecase "feature1開発" as 4
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 5
	// usecase "結合テスト" as 8
	// usecase "リリース" as 9
	// usecase "finish" as 10
	// usecase "feature2開発" as 6
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 7
	// usecase "結合テスト" as 8
	// usecase "リリース" as 9
	// usecase "finish" as 10

	// 1 --> 2
	// 2 --> 3
	// 3 --> 4
	// 4 --> 5
	// 5 --> 8
	// 8 --> 9
	// 9 --> 10
	// 3 --> 6
	// 6 --> 7
	// 7 --> 8

	// @enduml

}
