// Code generated from /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/Htql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package htql // Htql
import "github.com/antlr4-go/antlr/v4"

// BaseHtqlListener is a complete listener for a parse tree produced by HtqlParser.
type BaseHtqlListener struct{}

var _ HtqlListener = &BaseHtqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHtqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHtqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHtqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHtqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseHtqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseHtqlListener) ExitQuery(ctx *QueryContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseHtqlListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseHtqlListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseHtqlListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseHtqlListener) ExitElementList(ctx *ElementListContext) {}

// EnterFromStmt is called when production fromStmt is entered.
func (s *BaseHtqlListener) EnterFromStmt(ctx *FromStmtContext) {}

// ExitFromStmt is called when production fromStmt is exited.
func (s *BaseHtqlListener) ExitFromStmt(ctx *FromStmtContext) {}

// EnterWhereStmt is called when production whereStmt is entered.
func (s *BaseHtqlListener) EnterWhereStmt(ctx *WhereStmtContext) {}

// ExitWhereStmt is called when production whereStmt is exited.
func (s *BaseHtqlListener) ExitWhereStmt(ctx *WhereStmtContext) {}

// EnterConditionExpr is called when production conditionExpr is entered.
func (s *BaseHtqlListener) EnterConditionExpr(ctx *ConditionExprContext) {}

// ExitConditionExpr is called when production conditionExpr is exited.
func (s *BaseHtqlListener) ExitConditionExpr(ctx *ConditionExprContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseHtqlListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseHtqlListener) ExitCondition(ctx *ConditionContext) {}

// EnterAttributeExpr is called when production attributeExpr is entered.
func (s *BaseHtqlListener) EnterAttributeExpr(ctx *AttributeExprContext) {}

// ExitAttributeExpr is called when production attributeExpr is exited.
func (s *BaseHtqlListener) ExitAttributeExpr(ctx *AttributeExprContext) {}

// EnterAttributeNullCheck is called when production attributeNullCheck is entered.
func (s *BaseHtqlListener) EnterAttributeNullCheck(ctx *AttributeNullCheckContext) {}

// ExitAttributeNullCheck is called when production attributeNullCheck is exited.
func (s *BaseHtqlListener) ExitAttributeNullCheck(ctx *AttributeNullCheckContext) {}

// EnterLogicalOp is called when production logicalOp is entered.
func (s *BaseHtqlListener) EnterLogicalOp(ctx *LogicalOpContext) {}

// ExitLogicalOp is called when production logicalOp is exited.
func (s *BaseHtqlListener) ExitLogicalOp(ctx *LogicalOpContext) {}
