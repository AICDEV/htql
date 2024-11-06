# HTQL (Hyper Text Query Language)

HTQL is a SQL-like query language designed for extracting data from HTML structures. Its syntax is simple and intuitive, making it easy to use within other SQL adapters or applications.

## Table of Contents

- [Introduction](#introduction)
- [Usage Examples](#usage-examples)
    - [Basic Select](#basic-select)
    - [Select from Remote URL](#select-from-remote-url)
    - [Select with Binary Execution](#select-with-binary-execution)

## Introduction

HTQL allows users to extract structured data from HTML documents using familiar SQL syntax. With HTQL, you can select elements by type, apply filters, and even pull data from remote URLs, providing a powerful way to query HTML content in a standardized format.

## Usage Examples

### Basic Select

Use the `SELECT` statement to query HTML structures. Examples:

```sql
SELECT * FROM ./test.html                     -- Select all elements from a local file
SELECT p, div, h2 FROM ./test.html             -- Select specific elements (p, div, h2)
SELECT * FROM ./test.html WHERE attributes.class = 'title'
SELECT * FROM ./test.html WHERE attributes IS NOT NULL
SELECT * FROM ./test.html WHERE attributes.class = 'title' OR attributes.id = 'content'
SELECT span FROM ./test.html WHERE attributes.class = 'title' AND attributes.id = 'content'
SELECT span FROM ./test.html WHERE attributes.class = 'title' AND NOT attributes.id = 'content'
```

### Select from Remote URL

You can also query data directly from remote HTML documents by specifying the URL:
    
```sql
SELECT p, div, h2 FROM https://example.com
```