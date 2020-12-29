package gormgen

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
)

// The Parser is used to parse a directory and expose information about the structs defined in the files of this directory.
type Parser struct {
	dir         string
	pkg         *build.Package
	parsedFiles []*ast.File
}

// NewParser create a new parser instance.
// 实体类所在的目录
func NewParser(dir string) *Parser {
	return &Parser{
		dir: dir,
	}
}

// getPackage parse dir get go file and package
func (p *Parser) getPackage() {
	pkg, err := build.Default.ImportDir(p.dir, build.ImportComment)
	if err != nil {
		log.Fatalf("cannot process directory %s: %s", p.dir, err)
	}
	p.pkg = pkg
	fmt.Println("p.pkg: ", *pkg)

}

// parseGoFiles parse go file
func (p *Parser) parseGoFiles() {
	var parsedFiles []*ast.File
	fs := token.NewFileSet()
	for _, file := range p.pkg.GoFiles {
		file = p.dir + "/" + file
		parsedFile, err := parser.ParseFile(fs, file, nil, 0)
		if err != nil {
			log.Fatalf("parsing package: %s: %s\n", file, err)
		}
		fmt.Println("parsedFiles :", &parsedFile)
		parsedFiles = append(parsedFiles, parsedFile)
	}
	//fmt.Println("parsedFiles :", &parsedFiles)
	p.parsedFiles = parsedFiles
}

// parseTypes parse type of struct
func (p *Parser) parseTypes(file *ast.File) (ret []structConfig) {
	ast.Inspect(file, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}
		for _, spec := range decl.Specs {
			var (
				data structConfig
			)
			typeSpec, _ok := spec.(*ast.TypeSpec)
			if !_ok {
				continue
			}
			// We only care about struct declaration (for now)
			var structType *ast.StructType
			if structType, ok = typeSpec.Type.(*ast.StructType); !ok {
				continue
			}

			data.StructName = typeSpec.Name.Name
			data.StructNameLitle = strings.ToLower(data.StructName)
			for _, v := range structType.Fields.List {
				fmt.Println("struct V: ", v)
				fmt.Println("struct V.tag: ", v.Tag)
				fmt.Println("struct V.Names: ", v.Names)
				fmt.Println("struct V.Comment: ", v.Comment)
				fmt.Println("struct V.Type: ", v.Type)
				var (
					onlyField   fieldConfig
					optionField fieldConfig
				)
				// type is select expr, and sel is Model , else continue
				if _v, _ok := v.Type.(*ast.SelectorExpr); _ok {
					if _v.Sel.Name == "Model" {
						fmt.Println("======================")
						onlyField.FieldName = "ID"
						onlyField.FieldType = "uint"
						onlyField.ColumnName = gorm.ToDBName("ID")
						onlyField.HumpName = SQLColumnToHumpStyle(onlyField.ColumnName)
						data.OnlyFields = append(data.OnlyFields, onlyField)

						f1 := fieldConfig{}
						f1.FieldName = "CreatedAt"
						f1.FieldType = "time.Time"
						f1.ColumnName = gorm.ToDBName("CreatedAt")
						f1.HumpName = SQLColumnToHumpStyle(f1.ColumnName)
						data.OptionFields = append(data.OptionFields, f1)

						f2 := fieldConfig{}
						f2.FieldName = "UpdatedAt"
						f2.FieldType = "time.Time"
						f2.ColumnName = gorm.ToDBName("UpdatedAt")
						f2.HumpName = SQLColumnToHumpStyle(f2.ColumnName)
						data.OptionFields = append(data.OptionFields, f2)
					} else {
						optionField.FieldName = v.Names[0].String()
						optionField.ColumnName = gorm.ToDBName(optionField.FieldName)
						optionField.HumpName = SQLColumnToHumpStyle(optionField.ColumnName)
						if t, _ok := v.Type.(*ast.Ident); _ok {
							optionField.FieldType = t.String()
						}

						data.OptionFields = append(data.OptionFields, optionField)
					}
					continue
				}
				// get onlyField tag
				if v.Tag != nil {
					if strings.Contains(v.Tag.Value, "gorm") && strings.Contains(v.Tag.Value, "unique") ||
						strings.Contains(v.Tag.Value, "primary") {

						// type is ident, get onlyField type
						if t, _ok := v.Type.(*ast.Ident); _ok {
							onlyField.FieldType = t.String()
						}
						// get file name
						if len(v.Names) > 0 {
							onlyField.FieldName = v.Names[0].String()
							onlyField.ColumnName = gorm.ToDBName(onlyField.FieldName)
						}

						data.OnlyFields = append(data.OnlyFields, onlyField)
						//continue
					}
				}

				// type is ident, get onlyField type
				if t, _ok := v.Type.(*ast.Ident); _ok {
					optionField.FieldType = t.String()
				}
				// get file name

				if len(v.Names) > 0 {

					optionField.FieldName = v.Names[0].String()
					optionField.ColumnName = gorm.ToDBName(optionField.FieldName)
					optionField.HumpName = SQLColumnToHumpStyle(optionField.ColumnName)
				}

				data.OptionFields = append(data.OptionFields, optionField)
				fmt.Println("data onlyfild:", data.OnlyFields)

			}
			ret = append(ret, data)
		}
		return true
	})
	return
}

// Parse should be called before any type querying for the parser. It takes the directory to be parsed and extracts all the structs defined in this directory.
func (p *Parser) Parse() (ret []structConfig) {
	var (
		data []structConfig
	)

	p.getPackage()
	p.parseGoFiles()
	for _, f := range p.parsedFiles {
		data = append(data, p.parseTypes(f)...)
	}
	fmt.Println("data :", data)
	return data
}
