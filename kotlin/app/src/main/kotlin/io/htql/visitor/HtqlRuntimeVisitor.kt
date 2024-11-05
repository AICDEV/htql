package io.htql.visitor

import io.htql.model.HtqlNode
import io.htql.parser.HtqlParser
import io.htql.parser.HtqlVisitor
import org.antlr.v4.runtime.tree.ErrorNode
import org.antlr.v4.runtime.tree.ParseTree
import org.antlr.v4.runtime.tree.RuleNode
import org.antlr.v4.runtime.tree.TerminalNode
import org.jsoup.nodes.Document
import org.jsoup.nodes.Element
import java.util.*

class HtqlRuntimeVisitor(private val document: Document) : HtqlVisitor<List<HtqlNode>> {
    override fun visit(tree: ParseTree?): List<HtqlNode> {
        return tree?.accept(this) ?: emptyList()
    }

    override fun visitChildren(p0: RuleNode?): List<HtqlNode> {
        val result = mutableListOf<HtqlNode>()

        for (i in 0 until p0!!.childCount) {
            val childResult = p0.getChild(i).accept(this)
            result.addAll(childResult)
        }
        return result
    }

    override fun visitTerminal(node: TerminalNode?): List<HtqlNode> {
        return emptyList()
    }

    override fun visitErrorNode(p0: ErrorNode?): List<HtqlNode> {
        throw RuntimeException("Error parsing HTQL")
    }

    override fun visitQuery(ctx: HtqlParser.QueryContext?): List<HtqlNode> {

        var nodes = visit(ctx?.selectStmt())

        ctx?.whereStmt()?.let { whereCtx ->
            nodes = nodes.filter { node ->
                visitWhereStmt(whereCtx).contains(node)
            }
        }
        return nodes
    }

    override fun visitSelectStmt(ctx: HtqlParser.SelectStmtContext?): List<HtqlNode> {
        val elementsFilter = ctx?.elementList()?.IDENTIFIER() ?: emptyList()
        return if (elementsFilter.isEmpty()) {
            document.allElements.map { elementToHtqlNode(it) }
        } else {
            document.select(elementsFilter.joinToString(",")).map { elementToHtqlNode(it) }
        }
    }

    override fun visitElementList(ctx: HtqlParser.ElementListContext?): List<HtqlNode> {
        return emptyList()
    }

    override fun visitFromStmt(ctx: HtqlParser.FromStmtContext?): List<HtqlNode> {
        return emptyList()
    }

    override fun visitWhereStmt(ctx: HtqlParser.WhereStmtContext?): List<HtqlNode> {
        return visit(
            ctx?.conditionExpr()
        )
    }

    override fun visitConditionExpr(ctx: HtqlParser.ConditionExprContext?): List<HtqlNode> {
        var result = visitCondition(ctx?.condition(0))
        for (i in 1 until (ctx?.condition()?.size ?: 0)) {
            val logicalOpCtx = ctx?.logicalOp(i - 1)
            val rightCondition = visitCondition(ctx?.condition(i))
            result = applyLogicalOp(logicalOpCtx, result, rightCondition)
        }
        return result
    }

    override fun visitCondition(ctx: HtqlParser.ConditionContext?): List<HtqlNode> {
        val attributeExpr = ctx?.attributeExpr()?.IDENTIFIER() ?: ""
        val targetValue = ctx?.attributeExpr()?.STRING()?.text?.removeSurrounding("'")

        return document.select("[$attributeExpr='$targetValue']").map { elementToHtqlNode(it) }
    }

    override fun visitAttributeExpr(ctx: HtqlParser.AttributeExprContext?): List<HtqlNode> {
        return emptyList()
    }

    override fun visitAttributeNullCheck(ctx: HtqlParser.AttributeNullCheckContext?): List<HtqlNode> {
        return emptyList()
    }

    override fun visitLogicalOp(ctx: HtqlParser.LogicalOpContext?): List<HtqlNode> {
        return emptyList()
    }

    private fun elementToHtqlNode(element: Element): HtqlNode {
        return HtqlNode(
            type = element.tagName(),
            attributes = element.attributes().associate { it.key to it.value },
            children = element.children().map { elementToHtqlNode(it) },
            innerText = element.text()
        )
    }

    private fun applyLogicalOp(
        logicalOpCtx: HtqlParser.LogicalOpContext?,
        left: List<HtqlNode>,
        right: List<HtqlNode>
    ): List<HtqlNode> {
        return when (logicalOpCtx?.text?.uppercase(Locale.getDefault())) {
            "AND" -> left.intersect(right).toList()
            "OR" -> left.union(right).toList()
            else -> emptyList()
        }
    }
}