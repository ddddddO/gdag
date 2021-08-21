package gdag_test

import (
	g "github.com/ddddddO/gdag"
)

func Example() {
	goal := g.NewGoal("ゴール(目的)")

	design := g.NewTask("設計")
	review_design := g.NewTask("レビュー対応")

	develop_feature_1 := g.NewTask("feature1開発")
	develop_feature_1.AddNote("xxが担当")
	review_develop_feature_1 := g.NewTask("レビュー対応")

	develop_feature_2 := g.NewTask("feature2開発")
	develop_feature_2.AddNote("yyが担当")
	review_develop_feature_2 := g.NewTask("レビュー対応")

	test := g.NewTask("結合テスト")
	release := g.NewTask("リリース")
	finish := g.NewTask("finish")

	goal.Connect(design).Connect(review_design).Connect(develop_feature_1).Connect(review_develop_feature_1).Connect(test)
	review_design.Connect(develop_feature_2).Connect(review_develop_feature_2).Connect(test)
	test.Connect(release).Connect(finish)

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
