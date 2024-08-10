package checkofficeid

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "mylinter",
	Doc:  "checks if OfficeID is present when security tag is `personal_data`",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if structType, ok := n.(*ast.StructType); ok {
				checkStruct(pass, structType)
			}
			return true
		})
	}
	return nil, nil
}

func isInt(kind types.BasicKind) bool {
	switch kind {
	case types.Int, types.Int32, types.Int64, types.Uint, types.Uint32, types.Uint64:
		return true
	default:
		return false
	}
}

func checkStruct(pass *analysis.Pass, structType *ast.StructType) {
	hasPersonalDataTag := false
	hasOfficeID := false

	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			continue
		}
		fieldName := field.Names[0].Name

		if field.Tag != nil {
			tag := field.Tag.Value
			if tag == "`security:\"personal_data\"`" {
				hasPersonalDataTag = true
			}
		}

		if fieldName == "OfficeID" {
			if typ, ok := pass.TypesInfo.TypeOf(field.Type).(*types.Basic); ok && isInt(typ.Kind()) {
				hasOfficeID = true
			}
		}
	}

	if hasPersonalDataTag && !hasOfficeID {
		pass.Reportf(structType.Pos(), "struct with `personal_data` tag must have `OfficeID` of type uint")
	}
}
