package gosqltemplate

import (
	"embed"
	"errors"
	"path/filepath"
	"strings"
)

// 存储 SQL ID 和 SQL 文本的映射
var sqlTemplates = make(map[string]string)

// 加载main.sql
func InitFromEmbedFS(fs embed.FS, mainFilePath string) error {
	reader := &EmbeddedFileReader{FS: fs}
	return parseSQLFile(reader, mainFilePath)
}

func InitFromOS(mainFilePath string) error {
	reader := &OSFileReader{}
	return parseSQLFile(reader, mainFilePath)
}

// 读取文件内容
// 解析特定的语法，如 '--# ' 和 '--@'
// 更新 SQL ID 和语句的映射关系
// 导入和包含其他 SQL 文件
func parseSQLFile(fileReader FileReader, filePath string) error {
	lines, err := fileReader.ReadFile(filePath)
	if err != nil {
		return err
	}

	var currentID string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "--#") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				currentID = parts[1]
				sqlTemplates[currentID] = ""
			}
		} else if strings.HasPrefix(line, "--@") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				includedFilePath := filepath.Dir(filePath) + "/" + parts[1]
				err := parseSQLFile(fileReader, includedFilePath)
				if err != nil {
					return err
				}
			}
		} else if currentID != "" {
			sqlTemplates[currentID] += line + "\n"
		}
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
