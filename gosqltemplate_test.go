package gosqltemplate

import (
  "embed"
  "fmt"
  "testing"
)

//go:embed sql-templates/*
var sqlFiles embed.FS

func TestGetSql(T *testing.T) {
  err := Init(sqlFiles, "sql-templates/main.sql")
  if err != nil {
    panic(err)
  }
  userSql, err := Get("users.selectAll") // SELECT * FROM users;
  if err != nil {
    panic(err)
  }

  fmt.Println(userSql)
  postsSql, err := Get("posts.selectAll") // SELECT * FROM posts;

  if err != nil {
    panic(err)
  }

  fmt.Println(postsSql)
}
