# GoSQLTemplate

## 简介

在 Go 语言环境中，管理多个 SQL 文件可能相当复杂。GoSQLTemplate 是一个工具，旨在简化和优化 SQL 文件的管理过程。

## 快速入门

### 初始化和获取 SQL 语句

通过以下示例，您可以快速了解如何使用 SQLTemplate 管理您的 SQL 模板：

```go
gosqltemplate.Init("sql-templates/main.sql")
var sql, err := gosqltemplate.Get("sqlId")
```

### SQL 文件结构

项目的目录结构可能如下所示：

```
├── sql-templates
│   ├── main.sql
│   ├── posts.sql
│   ├── users.sql
```

#### 主要 SQL 文件 (main.sql)

```sql
--@ users.sql
--@ posts.sql
```

#### 用户 SQL 文件 (users.sql)

```sql
--# users.selectAll
select * from users
```

#### 帖子 SQL 文件 (posts.sql)

```sql
--# posts.selectAll
select * from posts
```

### 获取 SQL 语句

初始化main.sql，可以通过 SQL ID 来获取特定的 SQL 语句：

```go
gosqltemplate.Init("sql-templates/main.sql")
var usersSql,err := gosqltemplate.Get("users.selectAll")  // SELECT * FROM users;
var postsSql,err := gosqltemplate.Get("posts.selectAll")  // SELECT * FROM posts;
```

## 常用指令

- `--# {sqlId}`：定义一个 SQL ID，注意，ID 命名中间使用空格分割。
- `--@ {file.sql}`：导入其他 SQL 文件，注意，文件名中间使用空格分割。