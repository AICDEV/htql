// Code generated from /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/Htql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package htql // Htql
import "github.com/antlr4-go/antlr/v4"

type BaseHtqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHtqlVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitElementList(ctx *ElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitFromStmt(ctx *FromStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitWhereStmt(ctx *WhereStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitConditionExpr(ctx *ConditionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitAttributeExpr(ctx *AttributeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitAttributeNullCheck(ctx *AttributeNullCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHtqlVisitor) VisitLogicalOp(ctx *LogicalOpContext) interface{} {
	return v.VisitChildren(ctx)
}
