# HTQL

Language specification for HTQL (Hyper Text Query Language).

## Introduction

HTQL is a SQL like query language for extracting data from html structures. It is designed to be simple and easy to use
within other sql adapters.

## Usages

The syntax of HTQL is similar to SQL. The core data structure is a HtqlNode:

```kotlin
data class HtqlNode(
    val type: String,
    val attributes: Map<String, String>,
    val children: List<HtqlNode>,
    val innerText: String?,
)
```

Every result of a query is a list of HtqlNode or an empty list.

### Select

The select statement is used to query data from html structures. The syntax is as follows:

- SELECT * FROM document;
- SELECT p,div,h2 FROM document;
- SELECT * FROM document WHERE attributes.class = 'title';
- SELECT * FROM document WHERE attributes IS NOT NULL;
- SELECT * FROM document WHERE attributes.class = 'title' OR attributes.id = 'content';
- SELECT span FROM document WHERE attributes.class = 'title' AND attributes.id = 'content';
- SELECT span FROM document WHERE attributes.class = 'title' AND NOT attributes.id = 'content';

### Select from remote

HTQL can be used to query data from remote html structures. The syntax is as follows:

- SELECT p,div,h2 FROM https://example.com;
