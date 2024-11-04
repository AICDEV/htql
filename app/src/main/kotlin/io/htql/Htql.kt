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


fun main() {
    println("welcome to HTQL - HYPER TEXT QUERY LANGUAGE")

    println("Please enter your HTQL query:")

    val test = readlnOrNull() ?: exitProcess(1)

    val urlPattern = Regex("(http|https)://[a-zA-Z0-9_.-]+(:[0-9]+)?(/[a-zA-Z0-9_.-/?&%#=]*)?")

    var document: Document
    val urlMatch = urlPattern.find(test)
    if (urlMatch != null) {
        val url = urlMatch.value
        document = Jsoup.connect(url).get()
    } else {
        document = Jsoup.parse(File("app/src/main/resources/test.html"))
    }

    val lexer = HtqlLexer(CharStreams.fromString(test))
    val tokens = CommonTokenStream(lexer)
    val parser = HtqlParser(tokens)
    val tree = parser.query()


    val visitor = HtqlRuntimeVisitor(document = document)
    val nodes = visitor.visit(tree)


    nodes.forEach {
        println(it.type)
    }
}
