// Code generated from /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/Htql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package htql // Htql
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by HtqlParser.
type HtqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HtqlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by HtqlParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by HtqlParser#elementList.
	VisitElementList(ctx *ElementListContext) interface{}

	// Visit a parse tree produced by HtqlParser#fromStmt.
	VisitFromStmt(ctx *FromStmtContext) interface{}

	// Visit a parse tree produced by HtqlParser#whereStmt.
	VisitWhereStmt(ctx *WhereStmtContext) interface{}

	// Visit a parse tree produced by HtqlParser#conditionExpr.
	VisitConditionExpr(ctx *ConditionExprContext) interface{}

	// Visit a parse tree produced by HtqlParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by HtqlParser#attributeExpr.
	VisitAttributeExpr(ctx *AttributeExprContext) interface{}

	// Visit a parse tree produced by HtqlParser#attributeNullCheck.
	VisitAttributeNullCheck(ctx *AttributeNullCheckContext) interface{}

	// Visit a parse tree produced by HtqlParser#logicalOp.
	VisitLogicalOp(ctx *LogicalOpContext) interface{}
}
