package transformer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/digeon-inc/royle/pipe"
	"gorm.io/gorm/schema"
)

type strutInfo struct {
	name       string
	fieldNames []string
}

func SortColumnByGormModelFile(namer schema.Namer, tables []pipe.Table, dirs []string) ([]pipe.Table, error) {

	filePaths := make([]string, 100)
	var err error

	for _, dir := range dirs {
		if filePaths, err = getFilePaths(dir, filePaths); err != nil {
			return nil, err
		}
	}

	for _, filePath := range filePaths {
		structInfoList, err := getStructInfo(filePath)

		if err != nil {
			return nil, err
		}
		if len(structInfoList) == 0 {
			// ファイル内にstructがない場合はログを出力して、ソートせずにスルーする
			continue
		}

		for _, structInfo := range structInfoList {
			var table *pipe.Table

			// TODO:線形探索なので遅い
			for i, t := range tables {
				if t.TableName != namer.TableName(structInfo.name) {
					continue
				}
				table = &tables[i]
			}

			if table == nil {
				continue
			}

			columnMap := make(map[string]pipe.Column)
			for _, column := range table.Columns {
				columnMap[column.ColumnName] = column
			}

			var reorderedColumns []pipe.Column
			for _, fieldName := range structInfo.fieldNames {
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
		}

	}

	return tables, nil
}

func getStructInfo(filepath string) ([]strutInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filepath, nil, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}

	strutInfoList := make([]strutInfo, 1)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if structType, ok := x.Type.(*ast.StructType); ok {
				var structInfo strutInfo
				structInfo.name = x.Name.Name
				for _, field := range structType.Fields.List {
					for _, name := range field.Names {
						structInfo.fieldNames = append(structInfo.fieldNames, name.Name)
					}
				}
			}
		}
		return true
	})
	return strutInfoList, nil
}

func getFilePaths(dir string, paths []string) ([]string, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			paths = append(paths, path)
		}
		return nil
	})
	return paths, err
}
