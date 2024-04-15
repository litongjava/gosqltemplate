package gosqltemplate

import (
  "bufio"
  "embed"
  "errors"
  "path/filepath"
  "strings"
)

// 存储 SQL ID 和 SQL 文本的映射
var sqlTemplates = make(map[string]string)

// 加载main.sql
func Init(sqlFiles embed.FS, mainFilePath string) error {
  return parseSQLFile(sqlFiles, mainFilePath)
}

// 读取文件内容
// 解析特定的语法，如 '--# ' 和 '--@'
// 更新 SQL ID 和语句的映射关系
// 导入和包含其他 SQL 文件
func parseSQLFile(sqlFiles embed.FS, filePath string) error {

  file, err := sqlFiles.Open(filePath)
  if err != nil {
    return err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  var currentID string
  for scanner.Scan() {
    line := scanner.Text()
    line = strings.TrimSpace(line)
    if strings.HasPrefix(line, "--#") {
      // Handle SQL ID definition
      parts := strings.Fields(line)
      if len(parts) > 1 {
        currentID = parts[1]
        sqlTemplates[currentID] = ""
      }
    } else if strings.HasPrefix(line, "--@") {
      // Handle SQL file import
      parts := strings.Fields(line)
      if len(parts) > 1 {
        includedFilePath := filepath.Dir(filePath) + "/" + parts[1]
        err := parseSQLFile(sqlFiles, includedFilePath)
        if err != nil {
          return err
        }
      }
    } else if currentID != "" {
      sqlTemplates[currentID] += line + "\n"
    }
  }
  if err := scanner.Err(); err != nil {
    return err
  }
  return nil
}

// 从存储中检索 SQL ID 对应的 SQL 语句
// 如果 SQL ID 不存在，返回错误
func Get(sqlId string) (string, error) {
  sql, ok := sqlTemplates[sqlId]
  if !ok {
    return "", errors.New("SQL ID not found")
  }
  return sql, nil
}

func GetAll() map[string]string {
  return sqlTemplates
}
