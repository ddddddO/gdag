package gdag_test

import (
	g "github.com/ddddddO/gdag"
)

func Example() {
	var goal *g.Node = g.Goal("ゴール(目的)")

	var design *g.Node = g.Task("設計")
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

func ExampleGenerateUML() {
	var goal *g.Node = g.Goal("ゴール(目的)")

	var design *g.Node = g.Task("設計")
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
	// rectangle "ゴール(目的)" as 12
	// usecase "設計" as 13 #DarkGray
	// usecase "レビュー対応" as 14 #DarkGray
	// usecase "feature1開発" as 15
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 16
	// usecase "結合テスト" as 20
	// usecase "リリース" as 21
	// usecase "finish" as 22 #DarkGray
	// usecase "feature2開発" as 17 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 18
	// usecase "インフラ準備" as 19
	// note left
	// zzが担当
	// end note
	//
	// 12 --> 13
	// 13 --> 14
	// 14 --> 15
	// 15 --> 16
	// 16 --> 20
	// 20 --> 21
	// 21 --> 22
	// 14 --> 17
	// 17 --> 18
	// 18 --> 20
	// 14 --> 19
	// 19 --> 20
	//
	// @enduml
}

func ExampleGUML() {
	var goal *g.Node = g.G("ゴール(目的)")

	var design *g.Node = g.T("設計")
	reviewDesign := g.T("レビュー対応")

	developFeature1 := g.T("feature1開発")
	developFeature1.N("xxが担当")
	reviewDevelopFeature1 := g.T("レビュー対応")

	developFeature2 := g.T("feature2開発")
	developFeature2.N("yyが担当")
	reviewDevelopFeature2 := g.T("レビュー対応")

	prepareInfra := g.T("インフラ準備")
	prepareInfra.N("zzが担当")

	test := g.T("結合テスト")
	release := g.T("リリース")
	finish := g.T("finish")

	goal.C(design).C(reviewDesign).C(developFeature1).C(reviewDevelopFeature1).C(test)
	reviewDesign.C(developFeature2).C(reviewDevelopFeature2).C(test)
	reviewDesign.C(prepareInfra).C(test)
	test.C(release).C(finish)

	g.D(design, reviewDesign, developFeature2, finish)

	if err := g.GUML(goal); err != nil {
		panic(err)
	}
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

	g.Done(design, reviewDesign, developFeature2, finish)

	if err := g.GenerateCheckList(design); err != nil {
		panic(err)
	}
	// Output:
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

func ExampleGenerateGantt() {
	var goal *g.Node = g.Goal("ゴール(目的)")

	var design *g.Node = g.Task("設計")
	design.WithGanttStart("2021-9-3", 1)
	reviewDesign := g.Task("レビュー対応")
	reviewDesign.WithGantt(1)

	developFeature1 := g.Task("feature1開発")
	developFeature1.Note("xxが担当")
	developFeature1.WithGantt(1)
	reviewDevelopFeature1 := g.Task("レビュー対応")
	reviewDevelopFeature1.WithGantt(1)

	developFeature2 := g.Task("feature2開発")
	developFeature2.Note("yyが担当")
	developFeature2.WithGantt(4)
	reviewDevelopFeature2 := g.Task("レビュー対応")
	reviewDevelopFeature2.WithGantt(1)

	prepareInfra := g.Task("インフラ準備")
	prepareInfra.Note("zzが担当")
	prepareInfra.WithGantt(2)

	test := g.Task("結合テスト")
	test.WithGantt(1)
	release := g.Task("リリース")
	release.WithGantt(1)
	finish := g.Task("finish")
	finish.WithGantt(1)

	goal.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	if err := g.GenerateGantt(goal); err != nil {
		panic(err)
	}
	// Output:
	// <html>
	//     <body>
	//         <script src="https://cdn.jsdelivr.net/npm/mermaid/dist/mermaid.min.js"></script>
	//         <script>
	//             mermaid.initialize({ startOnLoad: true });
	//         </script>
	//
	//         Here is one mermaid diagram:
	//         <div class="mermaid">
	//             gantt
	//             dateFormat  YYYY-MM-DD
	//
	//             section ゴール(目的)
	// 	設計 :46,2021-9-3,1d
	// 	レビュー対応 :47,after 46,1d
	// 	feature1開発 :48,after 47,1d
	// 	レビュー対応 :49,after 48,1d
	// 	feature2開発 :50,after 47,4d
	// 	レビュー対応 :51,after 50,1d
	// 	インフラ準備 :52,after 47,2d
	// 	結合テスト :53,after 51,1d
	// 	リリース :54,after 53,1d
	// 	finish :55,after 54,1d
	//
	//         </div>
	//     </body>
	// </html>
}
