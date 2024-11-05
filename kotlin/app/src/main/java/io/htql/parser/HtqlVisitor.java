// Generated from Htql.g4 by ANTLR 4.13.2
package io.htql.parser;
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link HtqlParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface HtqlVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link HtqlParser#query}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitQuery(HtqlParser.QueryContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#selectStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSelectStmt(HtqlParser.SelectStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#elementList}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitElementList(HtqlParser.ElementListContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#fromStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFromStmt(HtqlParser.FromStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#whereStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitWhereStmt(HtqlParser.WhereStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#conditionExpr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConditionExpr(HtqlParser.ConditionExprContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#condition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCondition(HtqlParser.ConditionContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#attributeExpr}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeExpr(HtqlParser.AttributeExprContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#attributeNullCheck}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAttributeNullCheck(HtqlParser.AttributeNullCheckContext ctx);
	/**
	 * Visit a parse tree produced by {@link HtqlParser#logicalOp}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLogicalOp(HtqlParser.LogicalOpContext ctx);
}