package gdag_test

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func ExampleNode_UML() {
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
	// Unordered output:
	// @startuml
	// rectangle "ゴール(目的)" as 58
	// usecase "設計" as 59 #DarkGray
	// usecase "レビュー対応" as 60 #DarkGray
	// usecase "feature1開発" as 61
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 62
	// usecase "結合テスト" as 66
	// usecase "リリース" as 67
	// usecase "finish" as 68 #DarkGray
	// usecase "feature2開発" as 63 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 64
	// usecase "インフラ準備" as 65
	// note left
	// zzが担当
	// end note
	//
	// 58 --> 59
	// 59 --> 60
	// 60 --> 61
	// 61 --> 62
	// 62 --> 66
	// 66 --> 67
	// 67 --> 68
	// 60 --> 63
	// 63 --> 64
	// 64 --> 66
	// 60 --> 65
	// 65 --> 66
	//
	// @enduml
}

// nolint:govet
func ExampleNode_UML_ShortMethod() {
	var dag *g.Node = g.DAG("ゴール(目的)")

	var design *g.Node = g.T("設計")
	reviewDesign := g.T("レビュー対応")

	developFeature1 := g.T("feature1開発")
	developFeature1.N("xxが担当")
	reviewDevelopFeature1 := g.T("レビュー対応")

	developFeature2 := g.T("feature2開発").N("yyが担当")
	reviewDevelopFeature2 := g.T("レビュー対応")

	prepareInfra := g.T("インフラ準備").N("zzが担当")

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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
	// Unordered output:
	// @startuml
	// rectangle "ゴール(目的)" as 69
	// usecase "設計" as 70 #DarkGray
	// usecase "レビュー対応" as 71 #DarkGray
	// usecase "feature1開発" as 72
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 73
	// usecase "結合テスト" as 77
	// usecase "リリース" as 78
	// usecase "finish" as 79 #DarkGray
	// usecase "feature2開発" as 74 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 75
	// usecase "インフラ準備" as 76
	// note left
	// zzが担当
	// end note
	//
	// 69 --> 70
	// 70 --> 71
	// 71 --> 72
	// 72 --> 73
	// 73 --> 77
	// 77 --> 78
	// 78 --> 79
	// 71 --> 74
	// 74 --> 75
	// 75 --> 77
	// 71 --> 76
	// 76 --> 77
	//
	// @enduml
}

// nolint:govet
func ExampleNode_UML_Multiple() {
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(dagUML)

	fmt.Println()

	infraUML, err := prepareInfra.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(infraUML)
	// Unordered output:
	// @startuml
	// rectangle "ゴール(目的)" as 80
	// usecase "設計" as 81 #DarkGray
	// usecase "レビュー対応" as 82 #DarkGray
	// usecase "feature1開発" as 83
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 84
	// usecase "結合テスト" as 88
	// usecase "リリース" as 89
	// usecase "finish" as 90 #DarkGray
	// usecase "feature2開発" as 85 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 86
	// usecase "インフラ準備" as 87
	// note left
	// zzが担当
	// end note
	//
	// 80 --> 81
	// 81 --> 82
	// 82 --> 83
	// 83 --> 84
	// 84 --> 88
	// 88 --> 89
	// 89 --> 90
	// 82 --> 85
	// 85 --> 86
	// 86 --> 88
	// 82 --> 87
	// 87 --> 88
	//
	// @enduml
	//
	// @startuml
	// usecase "インフラ準備" as 87
	// note left
	// zzが担当
	// end note
	// usecase "結合テスト" as 88
	// usecase "リリース" as 89
	// usecase "finish" as 90 #DarkGray
	//
	// 87 --> 88
	// 88 --> 89
	// 89 --> 90
	//
	// @enduml
}
