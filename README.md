# gdag
Easily manage DAG with Go

[![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gdag.svg)](https://pkg.go.dev/github.com/ddddddO/gdag) [![GitHub release](https://img.shields.io/github/release/ddddddO/gdag.svg)](https://github.com/ddddddO/gdag/releases)

# Demo
## PlantUML

1. `go run main.go > tmp.pu`

```go
package main

import (
	g "github.com/ddddddO/gdag"
)

func main() {
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
}
```

```
@startuml
rectangle "ゴール(目的)" as 1
usecase "設計" as 2 #DarkGray
usecase "レビュー対応" as 3 #DarkGray
usecase "feature1開発" as 4
note left
xxが担当
end note
usecase "レビュー対応" as 5
usecase "結合テスト" as 9
usecase "リリース" as 10
usecase "finish" as 11 #DarkGray
usecase "feature2開発" as 6 #DarkGray
note left
yyが担当
end note
usecase "レビュー対応" as 7
usecase "インフラ準備" as 8
note left
zzが担当
end note

1 --> 2
2 --> 3
3 --> 4
4 --> 5
5 --> 9
9 --> 10
10 --> 11
3 --> 6
6 --> 7
7 --> 9
3 --> 8
8 --> 9

@enduml
```

2. tmp.pu to png

![image](https://github.com/ddddddO/gdag/blob/main/dag.png)


## CheckList

1. `go run main.go`

```go
package main

import (
	g "github.com/ddddddO/gdag"
)

func main() {
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
}
```

```
- [x] 設計
- [x] レビュー対応
- [ ] feature1開発
- [ ] レビュー対応
- [x] feature2開発
- [ ] レビュー対応
- [ ] インフラ準備
- [ ] 結合テスト
- [ ] リリース
- [x] finish
```

2. share with members
- [x] 設計
- [x] レビュー対応
- [ ] feature1開発
- [ ] レビュー対応
- [x] feature2開発
- [ ] レビュー対応
- [ ] インフラ準備
- [ ] 結合テスト
- [ ] リリース
- [x] finish

## etc
### short name version

```go
package main

import (
	g "github.com/ddddddO/gdag"
)

func main() {
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
}
```


# Reference
- [about DAG](https://nave-kazu.hatenablog.com/entry/2015/11/30/154810)
