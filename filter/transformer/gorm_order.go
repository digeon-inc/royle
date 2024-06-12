package transformer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/digeon-inc/royle/pipe"
	"gorm.io/gorm/schema"
)

func SortColumnByGormModelFile(namer schema.Namer, tables []pipe.Table, dirs []string) ([]pipe.Table, error) {

	paths := make(map[string]string)
	var err error

	for _, dir := range dirs {
		if paths, err = getFilePaths(dir, paths); err != nil {
			return nil, err
		}
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var resultErr error

	for i := range tables {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			table := &tables[i]
			filePath, ok := paths[table.TableName]
			if !ok {
				// テーブルのファイルがない場合はログを出力して、ソートせずにスルーする
				log.Printf("No matching file found for table: %s\n", table.TableName)
				return
			}

			content, err := os.ReadFile(filePath)
			if err != nil {
				mu.Lock()
				resultErr = err
				mu.Unlock()
				return
			}

			fieldNames, err := parseStructFields(namer, string(content))
			if err != nil {
				// ファイル内にstructがない(modelが宣言されてない)場合はログを出力して、ソートせずにスルーする
				log.Printf("Error parsing struct for table %s: %s\n", table.TableName, err.Error())
				return
			}

			columnMap := make(map[string]pipe.Column)
			for _, column := range table.Columns {
				columnMap[column.ColumnName] = column
			}

			var reorderedColumns []pipe.Column
			for _, fieldName := range fieldNames {
				if column, ok := columnMap[fieldName]; ok {
					reorderedColumns = append(reorderedColumns, column)
				}
			}

			// mysqlのデータベース内だけに存在する(ファイルに書かれてないカラム)は最後に追加する。
			ExistReorderedMap := make(map[string]bool)
			for _, column := range reorderedColumns {
				ExistReorderedMap[column.ColumnName] = true
			}
			for _, column := range table.Columns {
				if _, exist := ExistReorderedMap[column.ColumnName]; !exist {
					reorderedColumns = append(reorderedColumns, column)
				}
			}

			table.Columns = reorderedColumns
		}(i)
	}

	wg.Wait()

	return tables, resultErr
}

func parseStructFields(namer schema.Namer, fileContent string) ([]string, error) {
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
			// gormのcodeをみるかぎりColumnNameの第一引数は使われてない。https://github.com/go-gorm/gorm/blob/73a988ceb22651e01c968a9ec20ae1709e73c8e6/schema/naming.go#L61
			fields = append(fields, namer.ColumnName("", match[1]))
		}
	}

	return fields, nil
}

func getFilePaths(dir string, paths map[string]string) (map[string]string, error) {
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
