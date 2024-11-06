package parser

import (
	"fmt"
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

	nodes := v.Visit(ctx.SelectStmt()).([]HtqlNode)

	if whereCtx := ctx.WhereStmt(); whereCtx != nil {
		whereNodes := v.VisitWhereStmt(whereCtx.(*htql.WhereStmtContext))
		nodeMap := make(map[string]map[string]string)

		if whereNodes == nil {
			return []HtqlNode{}
		}

		for _, node := range whereNodes.([]HtqlNode) {
			nodeMap[node.Type] = node.Attributes
		}

		var filteredNodes []HtqlNode
		for _, node := range nodes {
			if attrs, exists := nodeMap[node.Type]; exists {
				match := true
				for key, value := range attrs {
					if node.Attributes[key] != value {
						match = false
						break
					}
				}
				if match {
					filteredNodes = append(filteredNodes, node)
				}
			}
		}
		nodes = filteredNodes
	}

	return nodes
}

func (v *HtqlRuntimeVisitor) VisitSelectStmt(ctx *htql.SelectStmtContext) interface{} {

	var nodes []HtqlNode

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

func (v *HtqlRuntimeVisitor) VisitConditionExpr(ctx *htql.ConditionExprContext) interface{} {

	result := v.VisitCondition(ctx.Condition(0).(*htql.ConditionContext)).([]HtqlNode)

	for i := 1; i < ctx.GetChildCount(); i++ {
		if condition, ok := ctx.GetChild(i).(*htql.ConditionContext); ok {

			logicalOperationContext := ctx.LogicalOp(i - 1).(*htql.LogicalOpContext)
			innerResult := v.VisitCondition(condition)

			result = v.applyLogicalOp(logicalOperationContext, result, innerResult)
		}
	}

	return result
}

func (v *HtqlRuntimeVisitor) VisitCondition(ctx *htql.ConditionContext) interface{} {
	attributeExpr := ctx.AttributeExpr().IDENTIFIER().GetText()
	valueExpr := ctx.AttributeExpr().STRING().GetText()

	var nodes []HtqlNode

	selector := fmt.Sprintf("[%s=%s]", attributeExpr, valueExpr)

	v.document.Find(selector).Each(func(i int, element *goquery.Selection) {
		nodes = append(nodes, v.elementToHtqlNode(element))
	})

	return nodes
}

func (v *HtqlRuntimeVisitor) VisitElementList(ctx *htql.ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *HtqlRuntimeVisitor) VisitLogicalOp(ctx *htql.LogicalOpContext) interface{} {
	return nil
}

func (v *HtqlRuntimeVisitor) applyLogicalOp(context *htql.LogicalOpContext, result []HtqlNode, result2 interface{}) []HtqlNode {
	return result
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
