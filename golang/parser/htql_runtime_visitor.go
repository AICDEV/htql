package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/antlr4-go/antlr/v4"
	htql "htql/htql_parser_go"
)

type HtqlNode struct {
	Type       string
	Attributes map[string]string
	Children   []HtqlNode
	InnerText  string
}

type HtqlRuntimeVisitor struct {
	*htql.BaseHtqlVisitor
	document *goquery.Document
}

func NewHtqlRuntimeVisitor(document *goquery.Document) *HtqlRuntimeVisitor {
	return &HtqlRuntimeVisitor{
		BaseHtqlVisitor: &htql.BaseHtqlVisitor{},
		document:        document,
	}
}

func (v *HtqlRuntimeVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return []HtqlNode{}
	}
	return tree.Accept(v)
}

func (v *HtqlRuntimeVisitor) VisitQuery(ctx *htql.QueryContext) interface{} {

	nodes := v.Visit(ctx.SelectStmt())

	return nodes
}

func (v *HtqlRuntimeVisitor) VisitSelectStmt(ctx *htql.SelectStmtContext) interface{} {

	nodes := []HtqlNode{}

	if ctx.ElementList() != nil {

		for i := 0; i < ctx.ElementList().GetChildCount(); i++ {
			el := ctx.ElementList().IDENTIFIER(i)
			if el != nil {
				v.document.Find(el.GetText()).Each(func(i int, element *goquery.Selection) {
					nodes = append(nodes, v.elementToHtqlNode(element))
				})
			}
		}
	} else {
		v.document.Each(func(i int, element *goquery.Selection) {
			nodes = append(nodes, v.elementToHtqlNode(element))
		})
	}

	return nodes
}

func (v *HtqlRuntimeVisitor) VisitWhereStmt(ctx *htql.WhereStmtContext) interface{} {
	return v.Visit(ctx.ConditionExpr())
}

func (v *HtqlRuntimeVisitor) VisitElementList(ctx *htql.ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *HtqlRuntimeVisitor) elementToHtqlNode(element *goquery.Selection) HtqlNode {
	attributes := make(map[string]string)
	element.Each(func(i int, sel *goquery.Selection) {
		for _, attr := range sel.Nodes[0].Attr {
			attributes[attr.Key] = attr.Val
		}
	})

	children := []HtqlNode{}
	element.Children().Each(func(i int, sel *goquery.Selection) {
		children = append(children, v.elementToHtqlNode(sel))
	})

	return HtqlNode{
		Type:       goquery.NodeName(element),
		Attributes: attributes,
		Children:   children,
		InnerText:  element.Text(),
	}
}
