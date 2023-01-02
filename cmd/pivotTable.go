package main

import (
	"encoding/json"
	"errors"
	"syscall/js"

	"github.com/xuri/excelize/v2"
)

// AddPivotTable provides the method to add pivot table by given pivot table
// options. Note that the same fields can not in Columns, Rows and Filter
// fields at the same time.
func AddPivotTable(f *excelize.File) func(this js.Value, args []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		ret := map[string]interface{}{"error": nil}
		if err := prepareArgs(args, []argsRule{
			{types: []js.Type{js.TypeString}},
		}); err != nil {
			ret["error"] = err.Error()
			return js.ValueOf(ret)
		}
		var opt excelize.PivotTableOptions
		if err := json.Unmarshal([]byte(args[0].String()), &opt); err != nil {
			ret["error"] = err.Error()
			return js.ValueOf(ret)
		}
		if err := f.AddPivotTable(&opt); err != nil {
			ret["error"] = err.Error()
		}
		return js.ValueOf(ret)
	}
}

func getPivotTableOptions(args js.Value) (excelize.PivotTableOptions, error) {
	var opt excelize.PivotTableOptions

	dataRange := args.Get("dataRange")
	if TypeCheckOpt(dataRange, "string") {
		opt.DataRange = dataRange.String()
	} else {
		return opt, errors.New("dataRange argument must be string")
	}
	
	pivotTableRange := args.Get("pivotTableRange")
	if TypeCheckOpt(pivotTableRange, "string") {
		opt.PivotTableRange = pivotTableRange.String()
	} else {
		return opt, errors.New("pivotTableRange argument must be string")
	}

	rows := args.Get("rows")
	if TypeCheckOpt(rows, "array") {
		var pivotTableFields []excelize.PivotTableField

		for i := 0; i < rows.Length(); i++ {
			arg := rows.Index(i)
			pivotTableField, err := getPivotTableField(arg)
			if err != nil {
				return opt, err
			}
			pivotTableFields = append(pivotTableFields, pivotTableField)	
		}

		opt.Rows = pivotTableFields
	} else {
		return opt, errors.New("pivotTableRange argument must be string")
	}
	return opt, nil
}

func getPivotTableField(args js.Value) (excelize.PivotTableField, error) {
	var pivotTableField excelize.PivotTableField

	compact := args.Get("compact")
	if TypeCheckOpt(compact, "boolean") {
		pivotTableField.Compact = compact.Bool()
	} else {
		return pivotTableField, errors.New("compact argument must be boolean")
	}

	data := args.Get("data")
	if TypeCheckOpt(data, "string") {
		pivotTableField.Data = data.String()
	} else {
		return pivotTableField, errors.New("compact argument must be boolean")
	}

	name := args.Get("name")
	if TypeCheckOpt(name, "string") {
		pivotTableField.Name = name.String()
	} else {
		return pivotTableField, errors.New("data argument must be string")
	}

	outline := args.Get("outline")
	if TypeCheckOpt(outline, "boolean") {
		pivotTableField.Outline = outline.Bool()
	} else {
		return pivotTableField, errors.New("outline argument must be boolean")
	}

	subtotal := args.Get("subtotal")
	if TypeCheckOpt(subtotal, "string") {
		pivotTableField.Subtotal = outline.String()
	} else {
		return pivotTableField, errors.New("subtotal argument must be string")
	}

	defaultSubtotal := args.Get("defaultSubtotal")
	if TypeCheckOpt(defaultSubtotal, "boolean") {
		pivotTableField.DefaultSubtotal = defaultSubtotal.Bool()
	} else {
		return pivotTableField, errors.New("defaultSubtotal argument must be boolean")
	}

	return pivotTableField, nil
}