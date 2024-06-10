package transformer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/digeon-inc/royle/pipe"
)

func parseStructFields(fileContent string) ([]string, error) {
	structRe := regexp.MustCompile(`(?s)type\s+\w+\s+struct\s*\{(.*?)\}`)
	fieldRe := regexp.MustCompile(`(?m)^\s*(\w+)\s+\w+.*$`)

	structMatches := structRe.FindStringSubmatch(fileContent)
	if len(structMatches) < 2 {
		return nil, fmt.Errorf("no struct found")
	}

	fieldsPart := structMatches[1]
	fieldMatches := fieldRe.FindAllStringSubmatch(fieldsPart, -1)

	var fields []string
	for _, match := range fieldMatches {
		if len(match) > 1 {
			fields = append(fields, camelToSnake(match[1]))
		}
	}

	return fields, nil
}

func camelToSnake(s string) string {
	if s == "" {
		return s
	}

	delimiter := "_"
	sLen := len(s)
	var snake string
	for i, current := range s {
		if i > 0 && i+1 < sLen {
			if current >= 'A' && current <= 'Z' {
				next := s[i+1]
				prev := s[i-1]
				if (next >= 'a' && next <= 'z') || (prev >= 'a' && prev <= 'z') {
					snake += delimiter
				}
			}
		}
		snake += string(current)
	}

	snake = strings.ToLower(snake)
	return snake
}

func getFilePaths(dir string) (map[string]string, error) {
	paths := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			fileNameWithoutExt := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			paths[fileNameWithoutExt] = path
		}
		return nil
	})
	return paths, err
}

func SortColumnByGorm(tables []pipe.Table, dir string) ([]pipe.Table, error) {
	paths, err := getFilePaths(dir)
	if err != nil {
		return nil, err
	}

	for i, table := range tables {
		filePath, ok := paths[table.TableName]
		if !ok {
			// マッチするテーブルがない場合はログを出力してスルーする
			fmt.Printf("No matching file found for table: %s\n", table.TableName)
			continue
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		fieldNames, err := parseStructFields(string(content))
		if err != nil {
			return nil, err
		}

		columnMap := make(map[string]pipe.Column)
		for _, column := range table.Columns {
			columnMap[strings.ToLower(column.ColumnName)] = column
		}

		var reorderedColumns []pipe.Column
		for _, fieldName := range fieldNames {
			if column, ok := columnMap[fieldName]; ok {
				reorderedColumns = append(reorderedColumns, column)
			}
		}

		tables[i].Columns = reorderedColumns
	}

	return tables, nil
}
