package gdag_test

import (
	"fmt"

	g "github.com/ddddddO/gdag"
)

func Example() {
	var dag *g.Node = g.DAG("ゴール(目的)")

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

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	uml, err := dag.UML()
	if err != nil {
		panic(err)
	}
	fmt.Println(uml)
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

func ExampleUML() {
	var dag *g.Node = g.DAG("ゴール(目的)")

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

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	uml, err := dag.UML()
	if err != nil {
		panic(err)
	}
	fmt.Println(uml)
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

func ExampleShortMethod() {
	var dag *g.Node = g.DAG("ゴール(目的)")

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

	dag.C(design).C(reviewDesign).C(developFeature1).C(reviewDevelopFeature1).C(test)
	reviewDesign.C(developFeature2).C(reviewDevelopFeature2).C(test)
	reviewDesign.C(prepareInfra).C(test)
	test.C(release).C(finish)

	g.D(design, reviewDesign, developFeature2, finish)

	uml, err := dag.UML()
	if err != nil {
		panic(err)
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

func ExampleCheckList() {
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

	checkList, err := design.CheckList()
	if err != nil {
		panic(err)
	}
	fmt.Println(checkList)
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

func ExampleMultipleUMLs() {
	var dag *g.Node = g.DAG("ゴール(目的)")

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

	dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
	reviewDesign.Con(prepareInfra).Con(test)
	test.Con(release).Con(finish)

	g.Done(design, reviewDesign, developFeature2, finish)

	dagUML, err := dag.UML()
	if err != nil {
		panic(err)
	}
	fmt.Println(dagUML)

	fmt.Println()

	infraUML, err := prepareInfra.UML()
	if err != nil {
		panic(err)
	}
	fmt.Println(infraUML)
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as 45
	// usecase "設計" as 46 #DarkGray
	// usecase "レビュー対応" as 47 #DarkGray
	// usecase "feature1開発" as 48
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 49
	// usecase "結合テスト" as 53
	// usecase "リリース" as 54
	// usecase "finish" as 55 #DarkGray
	// usecase "feature2開発" as 50 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 51
	// usecase "インフラ準備" as 52
	// note left
	// zzが担当
	// end note
	//
	// 45 --> 46
	// 46 --> 47
	// 47 --> 48
	// 48 --> 49
	// 49 --> 53
	// 53 --> 54
	// 54 --> 55
	// 47 --> 50
	// 50 --> 51
	// 51 --> 53
	// 47 --> 52
	// 52 --> 53
	//
	// @enduml
	//
	// @startuml
	// usecase "インフラ準備" as 52
	// note left
	// zzが担当
	// end note
	// usecase "結合テスト" as 53
	// usecase "リリース" as 54
	// usecase "finish" as 55 #DarkGray
	//
	// 52 --> 53
	// 53 --> 54
	// 54 --> 55
	//
	// @enduml
}

func ExampleMultipleCheckLists() {
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

	designCheckList, err := design.CheckList()
	if err != nil {
		panic(err)
	}
	fmt.Println(designCheckList)

	infraCheckList, err := prepareInfra.CheckList()
	if err != nil {
		panic(err)
	}
	fmt.Println(infraCheckList)
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
	//
	// - [ ] インフラ準備
	// - [ ] 結合テスト
	// - [ ] リリース
	// - [x] finish
}
