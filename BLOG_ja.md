# はじめに
何らかゴールを目指すとき、そのゴールに辿り着くために何をしないといけないか？を洗い出すことはままあるかと思います。
そして、その洗い出した各タスクがどんな依存関係にあるかをパッと把握したい・共有したいという場面もあると思います。

そこでこの記事では、その各タスクを見渡すための鳥瞰図を作成・管理する助け(？)となる Go パッケージを紹介します。

# 比較
鳥瞰図を作成するためのサービス・ツールはいくつかあると思いますが、ここでは簡単に

- 素の PlantUML で書かれたもの
- https://github.com/ddddddO/gdag を使って書かれたもの（今回紹介したい Go パッケージ。最終的に PlantUML 形式を出力します）

を比較しながら紹介します。

まず、以下のような依存関係をもったゴールがあるとします。

![](https://storage.googleapis.com/zenn-user-upload/9dde7a419c8b-20240914.png)

これを PlantUML で表すと以下です。
```
@startuml
rectangle "ゴール" as dag
usecase "設計" as design
usecase "feature1開発" as developFeature1
usecase "Finish!" as finish

dag --> design
design --> developFeature1
developFeature1 --> finish

@enduml
```

gdag を使うと、
```go
package main

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func main() {
	dag := g.DAG("ゴール")
	design := g.Task("設計")
	developFeature1 := g.Task("feature1開発")
	finish := g.Task("Finish!")

	dag.Con(design).Con(developFeature1).Con(finish)

	uml, err := dag.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
}
```

`Task` 関数は各タスクを表し、それらを `Con` メソッドで繋げることができます。依存関係を直観的に記述できているかと思いますが、行数としては素の PlantUML の方が短いですね...!
では、もう少しタスクが増えた場合を見ていきます。

ゴールに向かう途中で、

- 設計のレビュー対応をしなければならないと気付いた
- feature2 の開発もスコープに含めたい
- 各 feature の開発後は必ずレビュー対応する

となったため、タスクを鳥瞰図に追加したいとします。

素の PlantUML は、
```
@startuml
rectangle "ゴール" as dag
usecase "設計" as design
usecase "レビュー対応" as reviewDesign
usecase "feature1開発" as developFeature1
usecase "レビュー対応" as reviewDevelopFeature1
usecase "feature2開発" as developFeature2
usecase "レビュー対応" as reviewDevelopFeature2
usecase "Finish!" as finish

dag --> design
design --> reviewDesign
reviewDesign --> developFeature1
developFeature1 --> reviewDevelopFeature1
reviewDevelopFeature1 --> finish
reviewDesign --> developFeature2
developFeature2 --> reviewDevelopFeature2
reviewDevelopFeature2 --> finish

@enduml
```

初期の PlantUML から修正するのが煩雑になっているかと思います。
gdag ではどうでしょうか。

```go
package main

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func main() {
	dag := g.DAG("ゴール")
	design := g.Task("設計")
	reviewDesign := g.Task("レビュー対応")
	developFeature1 := g.Task("feature1開発")
	reviewDevelopFeature1 := g.Task("レビュー対応")
	developFeature2 := g.Task("feature2開発")
	reviewDevelopFeature2 := g.Task("レビュー対応")
	finish := g.Task("Finish!")

        dag.Con(design).Con(reviewDesign).Con(developFeature1).Con(reviewDevelopFeature1).Con(finish)
	reviewDesign.Con(developFeature2).Con(reviewDevelopFeature2).Con(finish)

	uml, err := dag.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)
}
```

途中に追加したタスク（設計のレビュー対応）は、間に `.Con(reviewDesign)` を追加するだけでよく、
新たに追加された依存関係（設計レビュー対応 -> feature2開発 -> そのレビュー対応 -> Finish!）は、`.Con` を続けて記述すればよいため、楽かなと思っています。

結構いいのではないでしょうか？手前味噌であれですが（笑）

# さらに
## 「タスクにメモを足したいんだけど。例えば誰が担当するかを付記したい」
できます！

```go
developFeature1 := g.Task("feature1開発").Note("xxが担当")
```
`Note` メソッドで、「feature1開発」に対して「xxが担当」というメモを追加できます。
![](https://storage.googleapis.com/zenn-user-upload/1244a2ef001c-20240914.png)

## 「なんのタスクが終わったかわかりやすく表せない？」
できます！

```go
g.Done(design, reviewDesign, developFeature1)
```
こういう感じに、gdag に用意されている`Done`関数に終わったタスクを入れてあげればグレーで着色してくれます。
![](https://storage.googleapis.com/zenn-user-upload/8ac9aac3b6f5-20240914.png)

## 「GitHubでサクッと共有したいんだけど？」
できます！

```go
mermaid, err := dag.Mermaid()
```
`Mermaid` メソッドを呼び出してあげれば、Mermaid記法で表されたコードが出力されるので、それを GitHub に張り付ければささっと共有できます。
※`Note` メソッドは対応してません

```
graph TD
classDef doneColor fill:#868787
1("ゴール")
2(["設計"]):::doneColor
3(["レビュー対応"]):::doneColor
4(["feature1開発"]):::doneColor
5(["レビュー対応"])
8(["Finish!"])
6(["feature2開発"])
7(["レビュー対応"])

1 --> 2
2 --> 3
3 --> 4
4 --> 5
5 --> 8
3 --> 6
6 --> 7
7 --> 8
```

![](https://storage.googleapis.com/zenn-user-upload/a4cc9a8b098c-20240914.png)

# さいごに
https://github.com/ddddddO/gdag
を最後に使ってから久しく経っていました。最近大きめのタスク分割をして整理しないとまずそうということで、このパッケージを使う機会がきまして、久々でしたがすんなり書くことができて嬉しかったのでこの記事を書いてみました。
よければ使ってみてください～

（蛇足: 以下スクショは昔のメモ書きです。課題とコツが書いてありますね。こんな感じにメモしていたんだな～と懐かしくなりました。）
![](https://storage.googleapis.com/zenn-user-upload/41a33607630a-20240916.png)

# ほんとに最後に
よければぜひ！SRE / CRE も絶賛募集中です！（2024/11/18 現在）
https://herp.careers/v1/openlogi/requisition-groups/486b8b01-6cf9-4434-8601-381c9c092e0d