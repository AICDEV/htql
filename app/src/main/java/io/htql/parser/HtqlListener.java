// Generated from Htql.g4 by ANTLR 4.13.2
package io.htql.parser;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link HtqlParser}.
 */
public interface HtqlListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link HtqlParser#query}.
	 * @param ctx the parse tree
	 */
	void enterQuery(HtqlParser.QueryContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#query}.
	 * @param ctx the parse tree
	 */
	void exitQuery(HtqlParser.QueryContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#selectStmt}.
	 * @param ctx the parse tree
	 */
	void enterSelectStmt(HtqlParser.SelectStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#selectStmt}.
	 * @param ctx the parse tree
	 */
	void exitSelectStmt(HtqlParser.SelectStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#elementList}.
	 * @param ctx the parse tree
	 */
	void enterElementList(HtqlParser.ElementListContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#elementList}.
	 * @param ctx the parse tree
	 */
	void exitElementList(HtqlParser.ElementListContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#fromStmt}.
	 * @param ctx the parse tree
	 */
	void enterFromStmt(HtqlParser.FromStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#fromStmt}.
	 * @param ctx the parse tree
	 */
	void exitFromStmt(HtqlParser.FromStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#whereStmt}.
	 * @param ctx the parse tree
	 */
	void enterWhereStmt(HtqlParser.WhereStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#whereStmt}.
	 * @param ctx the parse tree
	 */
	void exitWhereStmt(HtqlParser.WhereStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#conditionExpr}.
	 * @param ctx the parse tree
	 */
	void enterConditionExpr(HtqlParser.ConditionExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#conditionExpr}.
	 * @param ctx the parse tree
	 */
	void exitConditionExpr(HtqlParser.ConditionExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#condition}.
	 * @param ctx the parse tree
	 */
	void enterCondition(HtqlParser.ConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#condition}.
	 * @param ctx the parse tree
	 */
	void exitCondition(HtqlParser.ConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#attributeExpr}.
	 * @param ctx the parse tree
	 */
	void enterAttributeExpr(HtqlParser.AttributeExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#attributeExpr}.
	 * @param ctx the parse tree
	 */
	void exitAttributeExpr(HtqlParser.AttributeExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#attributeNullCheck}.
	 * @param ctx the parse tree
	 */
	void enterAttributeNullCheck(HtqlParser.AttributeNullCheckContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#attributeNullCheck}.
	 * @param ctx the parse tree
	 */
	void exitAttributeNullCheck(HtqlParser.AttributeNullCheckContext ctx);
	/**
	 * Enter a parse tree produced by {@link HtqlParser#logicalOp}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOp(HtqlParser.LogicalOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link HtqlParser#logicalOp}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOp(HtqlParser.LogicalOpContext ctx);
}