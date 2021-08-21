# gdag
Easily manage DAG with Go

[![Go Reference](https://pkg.go.dev/badge/github.com/ddddddO/gdag.svg)](https://pkg.go.dev/github.com/ddddddO/gdag) [![GitHub release](https://img.shields.io/github/release/ddddddO/gdag.svg)](https://github.com/ddddddO/gdag/releases)

# Demo

1. `go run main.go > tmp.pu`

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

	test := g.Task("結合テスト")
	release := g.Task("リリース")
	finish := g.Task("finish")

	goal.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(test)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(test)
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
usecase "結合テスト" as 8
usecase "リリース" as 9
usecase "finish" as 10 #DarkGray
usecase "feature2開発" as 6 #DarkGray
note left
yyが担当
end note
usecase "レビュー対応" as 7

1 --> 2
2 --> 3
3 --> 4
4 --> 5
5 --> 8
8 --> 9
9 --> 10
3 --> 6
6 --> 7
7 --> 8

@enduml
```

2. tmp.pu to png

![image](https://github.com/ddddddO/gdag/blob/main/dag.png)
