package main

import "fmt"

type nodeType string

const (
	rectangle = nodeType("rectangle")
	usecase   = nodeType("usecase")
)

type node struct {
	nodeType nodeType
	text     string
	as       string // plantumlでの識別子。自動で生成したいけど、例えばa, b, ...とかにしちゃうと、他の人がplantumlを編集するとき辛くなる。あと、識別子なので重複する場合はエラーとする。
	note     string
	parents  []*node
	children []*node
}

func newNode(nodeType nodeType, text, as string) *node {
	return &node{
		nodeType: nodeType,
		text:     text,
		as:       as,
	}
}

func newGoal(text, as string) *node {
	return newNode(rectangle, text, as)
}

func newTask(text, as string) *node {
	return newNode(usecase, text, as)
}

func (upstream *node) connect(current *node) *node {
	// ここはもう少し考える必要がありそう。
	for _, child := range upstream.children {
		if current.as == child.as {
			//panic("duplicate error")
			return current
		}
	}

	upstream.children = append(upstream.children, current)
	current.parents = append(current.parents, upstream)
	return current
}

func (current *node) addNote(note string) {
	current.note = note
}

func main() {
	goal := newGoal("ゴール(目的)", "goal")

	design := newTask("設計", "design")
	review_design := newTask("レビュー対応", "review_design")

	develop_feature_1 := newTask("feature1開発", "develop_feature_1")
	develop_feature_1.addNote("xxが担当")
	review_develop_feature_1 := newTask("レビュー対応", "review_develop_feature_1")

	develop_feature_2 := newTask("feature2開発", "develop_feature_2")
	develop_feature_2.addNote("yyが担当")
	review_develop_feature_2 := newTask("レビュー対応", "review_develop_feature_2")

	test := newTask("結合テスト", "test")
	release := newTask("リリース", "release")
	finish := newTask("finish", "finish")

	goal.connect(design).connect(review_design).connect(develop_feature_1).connect(review_develop_feature_1).connect(test)
	review_design.connect(develop_feature_2).connect(review_develop_feature_2).connect(test)
	test.connect(release).connect(finish)

	if err := generateUML(goal); err != nil {
		panic(err)
	}
}

func generateUML(goal *node) error {
	fmt.Println("@startuml")

	fmt.Println(generateComponents(goal))
	fmt.Println(generateRelations(goal))

	fmt.Println("@enduml")
	return nil
}

func generateComponents(goal *node) string {
	generateComponent(goal)
	return dstComponents
}

var dstComponents string // 要リファクタ
func generateComponent(n *node) {
	if n.nodeType == rectangle {
		dstComponents += fmt.Sprintf("rectangle \"%s\" as %s\n", n.text, n.as)
	}
	if n.nodeType == usecase {
		dstComponents += fmt.Sprintf("usecase \"%s\" as %s\n", n.text, n.as)
	}
	if len(n.note) != 0 {
		dstComponents += fmt.Sprintf("note left\n%s\nend note\n", n.note)
	}

	for _, child := range n.children {
		generateComponent(child)
	}
}

var dstRelations string // 要リファクタ
func generateRelations(goal *node) string {
	generateRelation(goal)
	return dstRelations
}

// 重複する関係を除外する
func generateRelation(n *node) {
	r := fmt.Sprintf("%s --> ", n.as)
	for _, child := range n.children {
		dstRelations += r + child.as + "\n"
		generateRelation(child)
	}
}
