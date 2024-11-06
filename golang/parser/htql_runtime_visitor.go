package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antlr4-go/antlr/v4"
	htql "htql/htql_parser_go"
	"strings"
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
		whereNodes := v.VisitWhereStmt(whereCtx.(*htql.WhereStmtContext)).([]HtqlNode)
		nodeMap := make(map[string][]map[string]string)

		for _, node := range whereNodes {
			nodeMap[node.Type] = append(nodeMap[node.Type], node.Attributes)
		}

		var filteredNodes []HtqlNode
		for _, node := range nodes {
			if attrsList, exists := nodeMap[node.Type]; exists {
				match := false
				for _, attrs := range attrsList {
					matches := true
					for key, value := range attrs {
						if node.Attributes[key] != value {
							matches = false
							break
						}
					}
					if matches {
						match = true
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
		op := ctx.LogicalOp(i - 1)
		condition := ctx.Condition(i)
		if condition != nil {
			innerResult := v.VisitCondition(ctx.Condition(i).(*htql.ConditionContext)).([]HtqlNode)
			result = v.applyLogicalOp(op.(*htql.LogicalOpContext), result, innerResult)
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

func (v *HtqlRuntimeVisitor) applyLogicalOp(ctx *htql.LogicalOpContext, left []HtqlNode, right []HtqlNode) []HtqlNode {
	op := strings.ToUpper(ctx.GetText())

	switch op {
	case "AND":
		return intersect(left, right)
	case "OR":
		return union(left, right)
	default:
		return []HtqlNode{}
	}
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

func intersect(left []HtqlNode, right []HtqlNode) []HtqlNode {
	nodeMap := make(map[string]map[string]string)
	for _, node := range left {
		key := node.Type
		nodeMap[key] = node.Attributes
	}

	var result []HtqlNode
	for _, node := range right {
		key := node.Type
		if existingAttrs, exists := nodeMap[key]; exists && equalAttributes(existingAttrs, node.Attributes) {
			result = append(result, node)
		}
	}

	return result
}

func union(left []HtqlNode, right []HtqlNode) []HtqlNode {
	nodeMap := make(map[string]map[string]string)

	// Add all nodes from the left side to the map
	for _, node := range left {
		key := node.Type
		nodeMap[key] = node.Attributes
	}

	// Add nodes from the right side that are not already in the map
	for _, node := range right {
		key := node.Type
		// If the node type is not in the map or if the attributes differ, add it to the left slice
		if existingAttrs, exists := nodeMap[key]; !exists || !equalAttributes(existingAttrs, node.Attributes) {
			left = append(left, node)
			nodeMap[key] = node.Attributes
		}
	}

	return left
}

func equalAttributes(attrs1, attrs2 map[string]string) bool {
	if len(attrs1) != len(attrs2) {
		return false
	}

	for key, val1 := range attrs1 {
		if val2, exists := attrs2[key]; !exists || val1 != val2 {
			return false
		}
	}

	return true
}
