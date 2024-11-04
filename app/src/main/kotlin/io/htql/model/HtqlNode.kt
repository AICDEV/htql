package io.htql.model

data class HtqlNode(
    val type: String,
    val attributes: Map<String, String>,
    val children: List<HtqlNode>,
    val innerText: String?,
)
