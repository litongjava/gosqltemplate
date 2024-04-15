# GoSQLTemplate

## Introduction

In the Go language environment, managing multiple SQL files can be quite complex. GoSQLTemplate is a tool designed to simplify and optimize the management process of SQL files.

## Quick Start

### Initialization and Retrieval of SQL Statements

Through the following example, you can quickly understand how to manage your SQL templates using SQLTemplate:

```go
gosqltemplate.InitFromOS("../sql-templates/main.sql")
var sql, err := gosqltemplate.Get("sqlId")
```

### SQL File Structure

The project's directory structure might look like this:

```
├── sql-templates
│   ├── main.sql
│   ├── posts.sql
│   ├── users.sql
```

#### Main SQL file (main.sql)

```sql
--@ users.sql
--@ posts.sql
```

#### User SQL file (users.sql)

```sql
--# users.selectAll
select * from users
```

#### Posts SQL file (posts.sql)

```sql
--# posts.selectAll
select * from posts
```

### Retrieving SQL Statements

By initializing main.sql, you can retrieve specific SQL statements through the SQL ID:

```go
package services

import (
	"fmt"
	"github.com/litongjava/gosqltemplate"
	"testing"
)

func TestGetUserSql(t *testing.T) {
	err := gosqltemplate.InitFromOS("../sql-templates/main.sql")
	if err != nil {
		panic(err)
	}
	postSelectAll, err := gosqltemplate.Get("posts.selectAll")
	if err != nil {
		panic(err)
	}
	fmt.Println(postSelectAll)
}
```

## Common Commands

- `--# {sqlId}`: Defines an SQL ID, note that spaces are used to separate words in the ID naming.
- `--@ {file.sql}`: Imports other SQL files, note that spaces are used to separate words in the file name.

## support for embed.FS
embedfs.go

```go
import "embed"

//go:embed sql-templates/*
var SqlFiles embed.FS
```

main.go
```go
err := gosqltemplate.InitFromEmbedFS(SqlFiles, "sql-templates/main.sql")
if err != nil {
  panic(err)
}
```