package io.htql

import io.htql.parser.HtqlLexer
import io.htql.parser.HtqlParser
import io.htql.visitor.HtqlRuntimeVisitor
import org.antlr.v4.runtime.CharStreams
import org.antlr.v4.runtime.CommonTokenStream
import org.jsoup.Jsoup
import org.jsoup.nodes.Document
import java.io.File
import kotlin.system.exitProcess


val urlPattern = Regex("(http|https)://[a-zA-Z0-9_.-]+(:[0-9]+)?(/[a-zA-Z0-9_.-/?&%#=]*)?")

fun main() {
    println("welcome to HTQL - HYPER TEXT QUERY LANGUAGE")

    println("Please enter your HTQL query:")

    val expression = readlnOrNull() ?: exitProcess(1)

    evaluateDocument(expression) { document ->
        val lexer = HtqlLexer(CharStreams.fromString(expression))
        val tokens = CommonTokenStream(lexer)
        val parser = HtqlParser(tokens)
        val tree = parser.query()


        val visitor = HtqlRuntimeVisitor(document = document)
        val nodes = visitor.visit(tree)

        println(nodes)
    }
}

fun evaluateDocument(expression: String, block: (document: Document) -> Unit) {
    val urlMatch = urlPattern.find(expression)
    if (urlMatch != null) {
        val url = urlMatch.value
        block(Jsoup.connect(url).get())
    } else {
        block(Jsoup.parse(File("app/src/main/resources/test.html")))
    }
}
