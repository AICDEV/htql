#!/bin/sh

alias antlr4='java -Xmx500M -cp "/Users/aicdev/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -visitor -package htql /Users/aicdev/coding/htql/kotlin/app/src/main/antlr/*.g4 -o htql_parser_go