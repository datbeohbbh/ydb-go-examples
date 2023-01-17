# Basic example

Basic example demonstrates the possibilities of `YDB` with `xorm` ORM framework:
  - Create/Drop table with struct or struct pointer with `query_mode=scheme`
  - Insert multiple records by slice of pointer on one table with `query_mode=data`
  - Update/Delete records on table
  - Transaction with `engine.Transaction(...)`
  - Upsert/Replace multiple records by `[]map[string]interface{}` or fetch data from other table using `xorm.io/builder` package
  - Explain query mode for getting `AST` and `Plan` of query processing with `query_mode=explain`
  - Select with `query_mode=scan`