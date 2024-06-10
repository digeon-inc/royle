package transformer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/digeon-inc/royle/pipe"
)

// ParseStructFields parses the struct fields from a string and returns them in lowercase
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
			fields = append(fields, strings.ToLower(match[1]))
		}
	}

	return fields, nil
}

// FindMatchingFiles searches for files containing structs that match the given table names
func findMatchingFiles(tables []pipe.Table, dir string) (map[string]string, error) {
	matchingFiles := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			for _, table := range tables {
				structRe := regexp.MustCompile(fmt.Sprintf(`type\s+%s\s+struct`, table.TableName))
				if structRe.Match(content) {
					matchingFiles[table.TableName] = path
					break
				}
			}
		}
		return nil
	})

	// エラーが発生した場合はログを出力してスルーする
	if err != nil {
		fmt.Println("Error occurred while searching for matching files:", err)
	}

	return matchingFiles, nil
}

// ReorderColumns reorders the columns in each Table to match the order in the corresponding struct file
func SortColumnByGorm(tables []pipe.Table, dir string) ([]pipe.Table, error) {
	// Find matching files for each table
	matchingFiles, err := findMatchingFiles(tables, dir)
	if err != nil {
		return nil, err
	}
	// Iterate over each table and reorder its columns
	for i, table := range tables {
		filePath, ok := matchingFiles[table.TableName]
		if !ok {
			// マッチするテーブルがない場合はログを出力してスルーする
			fmt.Printf("No matching file found for table: %s\n", table.TableName)
			continue
		}

		// Read the file content
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		// Parse the struct fields from the file content
		fieldNames, err := parseStructFields(string(content))
		if err != nil {
			return nil, err
		}

		// Create a map of column names to columns for quick lookup
		columnMap := make(map[string]pipe.Column)
		for _, column := range table.Columns {
			columnMap[strings.ToLower(column.ColumnName)] = column
		}

		// Reorder the columns based on the field names
		var reorderedColumns []pipe.Column
		for _, fieldName := range fieldNames {
			if column, ok := columnMap[fieldName]; ok {
				reorderedColumns = append(reorderedColumns, column)
			}
		}

		// Update the table's columns with the reordered columns
		tables[i].Columns = reorderedColumns
	}

	return tables, nil
}
