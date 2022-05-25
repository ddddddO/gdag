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
}

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
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as 56
	// usecase "設計" as 57 #DarkGray
	// usecase "レビュー対応" as 58 #DarkGray
	// usecase "feature1開発" as 59
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 60
	// usecase "結合テスト" as 64
	// usecase "リリース" as 65
	// usecase "finish" as 66 #DarkGray
	// usecase "feature2開発" as 61 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 62
	// usecase "インフラ準備" as 63
	// note left
	// zzが担当
	// end note
	//
	// 56 --> 57
	// 57 --> 58
	// 58 --> 59
	// 59 --> 60
	// 60 --> 64
	// 64 --> 65
	// 65 --> 66
	// 58 --> 61
	// 61 --> 62
	// 62 --> 64
	// 58 --> 63
	// 63 --> 64
	//
	// @enduml
}

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
	// Output:
	// @startuml
	// rectangle "ゴール(目的)" as 67
	// usecase "設計" as 68 #DarkGray
	// usecase "レビュー対応" as 69 #DarkGray
	// usecase "feature1開発" as 70
	// note left
	// xxが担当
	// end note
	// usecase "レビュー対応" as 71
	// usecase "結合テスト" as 75
	// usecase "リリース" as 76
	// usecase "finish" as 77 #DarkGray
	// usecase "feature2開発" as 72 #DarkGray
	// note left
	// yyが担当
	// end note
	// usecase "レビュー対応" as 73
	// usecase "インフラ準備" as 74
	// note left
	// zzが担当
	// end note
	//
	// 67 --> 68
	// 68 --> 69
	// 69 --> 70
	// 70 --> 71
	// 71 --> 75
	// 75 --> 76
	// 76 --> 77
	// 69 --> 72
	// 72 --> 73
	// 73 --> 75
	// 69 --> 74
	// 74 --> 75
	//
	// @enduml
	//
	// @startuml
	// usecase "インフラ準備" as 74
	// note left
	// zzが担当
	// end note
	// usecase "結合テスト" as 75
	// usecase "リリース" as 76
	// usecase "finish" as 77 #DarkGray
	//
	// 74 --> 75
	// 75 --> 76
	// 76 --> 77
	//
	// @enduml
}
