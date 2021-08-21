package gdag_test

import (
	g "github.com/ddddddO/gdag"
)

func Example() {
	goal := g.NewGoal("ゴール(目的)", "goal")

	design := g.NewTask("設計", "design")
	review_design := g.NewTask("レビュー対応", "review_design")

	develop_feature_1 := g.NewTask("feature1開発", "develop_feature_1")
	develop_feature_1.AddNote("xxが担当")
	review_develop_feature_1 := g.NewTask("レビュー対応", "review_develop_feature_1")

	develop_feature_2 := g.NewTask("feature2開発", "develop_feature_2")
	develop_feature_2.AddNote("yyが担当")
	review_develop_feature_2 := g.NewTask("レビュー対応", "review_develop_feature_2")

	test := g.NewTask("結合テスト", "test")
	release := g.NewTask("リリース", "release")
	finish := g.NewTask("finish", "finish")

	goal.Connect(design).Connect(review_design).Connect(develop_feature_1).Connect(review_develop_feature_1).Connect(test)
	review_design.Connect(develop_feature_2).Connect(review_develop_feature_2).Connect(test)
	test.Connect(release).Connect(finish)

	if err := g.GenerateUML(goal); err != nil {
		panic(err)
	}
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as goal
	// usecase "設計" as design
	// usecase "レビュー対応" as review_design
	// usecase "feature1開発" as develop_feature_1
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as review_develop_feature_1
	// usecase "結合テスト" as test
	// usecase "リリース" as release
	// usecase "finish" as finish
	// usecase "feature2開発" as develop_feature_2
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as review_develop_feature_2
	// usecase "結合テスト" as test
	// usecase "リリース" as release
	// usecase "finish" as finish

	// goal --> design
	// design --> review_design
	// review_design --> develop_feature_1
	// develop_feature_1 --> review_develop_feature_1
	// review_develop_feature_1 --> test
	// test --> release
	// release --> finish
	// review_design --> develop_feature_2
	// develop_feature_2 --> review_develop_feature_2
	// review_develop_feature_2 --> test

	// @enduml
}
