// Code generated from /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/Htql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package htql // Htql
import "github.com/antlr4-go/antlr/v4"

// HtqlListener is a complete listener for a parse tree produced by HtqlParser.
type HtqlListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterFromStmt is called when entering the fromStmt production.
	EnterFromStmt(c *FromStmtContext)

	// EnterWhereStmt is called when entering the whereStmt production.
	EnterWhereStmt(c *WhereStmtContext)

	// EnterConditionExpr is called when entering the conditionExpr production.
	EnterConditionExpr(c *ConditionExprContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterAttributeExpr is called when entering the attributeExpr production.
	EnterAttributeExpr(c *AttributeExprContext)

	// EnterAttributeNullCheck is called when entering the attributeNullCheck production.
	EnterAttributeNullCheck(c *AttributeNullCheckContext)

	// EnterLogicalOp is called when entering the logicalOp production.
	EnterLogicalOp(c *LogicalOpContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitFromStmt is called when exiting the fromStmt production.
	ExitFromStmt(c *FromStmtContext)

	// ExitWhereStmt is called when exiting the whereStmt production.
	ExitWhereStmt(c *WhereStmtContext)

	// ExitConditionExpr is called when exiting the conditionExpr production.
	ExitConditionExpr(c *ConditionExprContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitAttributeExpr is called when exiting the attributeExpr production.
	ExitAttributeExpr(c *AttributeExprContext)

	// ExitAttributeNullCheck is called when exiting the attributeNullCheck production.
	ExitAttributeNullCheck(c *AttributeNullCheckContext)

	// ExitLogicalOp is called when exiting the logicalOp production.
	ExitLogicalOp(c *LogicalOpContext)
}
