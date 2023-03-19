package main

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func main() {
	flow := g.DAG("システム障害対応フロー")
	checkEvent := g.Task("イベントの確認").Note("システムエラーやユーザーからの申告。\nこの時点ではシステム障害ではない。")
	checkMatter := g.Task("検知・事象の確認").Note("システムが本来の機能を果たせていない\n可能性を検知した段階で障害対応を開始。")
	checkBusinessImpact := g.Task("業務影響調査")
	checkCause := g.Task("原因調査")
	recoveryResponse := g.Task("復旧対応")
	permanentResponse := g.Task("恒久対応")
	faultAnalysis := g.Task("障害分析")
	recurrencePreventionMeasures := g.Task("再発防止策")
	end := g.Task("終了")

	flow.Con(checkEvent).Con(checkMatter)

	checkMatter.Fanout(checkBusinessImpact, checkCause).Con(recoveryResponse)
	// Fanout method is same below.
	// checkMatter.Con(checkBusinessImpact).Con(recoveryResponse)
	// checkMatter.Con(checkCause).Con(recoveryResponse)

	recoveryResponse.Con(permanentResponse)
	permanentResponse.Con(faultAnalysis)
	permanentResponse.Con(recurrencePreventionMeasures)
	faultAnalysis.Con(end)
	recurrencePreventionMeasures.Con(end)

	uml, err := flow.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)

	// mermaid, err := flow.Mermaid()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	// fmt.Println(mermaid)
}