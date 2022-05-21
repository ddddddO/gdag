package main

import (
	"fmt"
	"os"

	g "github.com/ddddddO/gdag"
)

func main() {
	dag := g.DAG("夜ご飯")

	// みそ汁
	dagMisosiru, inputMiso := misosiru()
	// 生姜焼き
	dagSyougayaki, inputTare := syougayaki()
	// サラダ
	dagSalad, cutKyabetu := kyabetunosengiri()
	// だし巻き卵
	dagDasimaki, bakeTamago := dasimakitamago()

	dag.Con(dagMisosiru)
	dag.Con(dagSyougayaki)
	dag.Con(dagSalad)
	dag.Con(dagDasimaki)

	finish := g.Task("完成")
	inputMiso.Con(finish)
	inputTare.Con(finish)
	cutKyabetu.Con(finish)
	bakeTamago.Con(finish)

	uml, err := dag.UML()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(uml)

	fmt.Println()

	mermaid, err := dag.Mermaid()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(mermaid)

	fmt.Println()

	checklist, err := dag.CheckList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(checklist)
}

func misosiru() (*g.Node, *g.Node) {
	dagMisosiru := g.DAG("みそ汁")
	boilDasiziru := g.Task("だし汁を沸かす")
	inputWakame := g.Task("わかめを入れる")
	inputNegi := g.Task("ネギを入れる")
	stopBoil := g.Task("沸騰したら火を止める")
	inputMiso := g.Task("味噌を溶かす")
	dagMisosiru.Con(boilDasiziru)
	boilDasiziru.Con(inputWakame).Con(stopBoil)
	boilDasiziru.Con(inputNegi).Con(stopBoil)
	stopBoil.Con(inputMiso)

	return dagMisosiru, inputMiso
}

func syougayaki() (*g.Node, *g.Node) {
	dagSyougayaki := g.DAG("生姜焼き")

	dagTare := g.DAG("タレ")
	grateSyouga := g.Task("生姜ひとかけの3/4をおろす")
	inputSyouga := g.Task("残りの生姜を細切りにして入れる")

	inputSyoyu := g.Task("醤油大さじ5")
	inputMirin := g.Task("みりん大さじ3")
	inputSake := g.Task("酒大さじ1")
	inputSatou := g.Task("砂糖大さじ1")
	mergeTare := g.Task("混ぜる")
	dagTare.Con(grateSyouga).Con(inputSyouga)
	grateSyouga.Con(mergeTare)
	dagTare.Con(inputSyoyu).Con(mergeTare)
	dagTare.Con(inputMirin).Con(mergeTare)
	dagTare.Con(inputSake).Con(mergeTare)
	dagTare.Con(inputSatou).Con(mergeTare)

	dagPork := g.DAG("豚肉")
	bakePork := g.Task("薄く焼き目が付くまで焼く")
	inputTare := g.Task("タレと絡めて炒める")
	inputSyouga.Con(bakePork)
	dagPork.Con(bakePork).Con(inputTare)
	mergeTare.Con(inputTare)
	dagSyougayaki.Con(dagPork)
	dagSyougayaki.Con(dagTare)

	return dagSyougayaki, inputTare
}

func kyabetunosengiri() (*g.Node, *g.Node) {
	dagSalad := g.DAG("キャベツの千切り")
	cutKyabetu := g.Task("キャベツを切って盛り付ける")
	dagSalad.Con(cutKyabetu)
	return dagSalad, cutKyabetu
}

func dasimakitamago() (*g.Node, *g.Node) {
	dagDasimaki := g.DAG("だし巻き卵")
	mixTamago := g.Task("卵2個とつゆのだしを混ぜる").Note("つゆのだしは4,5滴でいいかも")
	bakeTamago := g.Task("たまごを焼く")
	dagDasimaki.Con(mixTamago).Con(bakeTamago)
	return dagDasimaki, bakeTamago
}
