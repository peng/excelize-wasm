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
		fields, err := getPivotTableFields(rows)
		if err != nil {
			return opt, err
		}
		opt.Rows = fields
	} else {
		return opt, errors.New("rows argument must be array")
	}

	columns := args.Get("columns")
	if TypeCheckOpt(columns, "array") {
		fields, err := getPivotTableFields(columns)
		if err != nil {
			return opt, err
		}
		opt.Columns = fields
	} else {
		return opt, errors.New("columns argument must be array")
	}

	data := args.Get("data")
	if TypeCheckOpt(data, "array") {
		fields, err := getPivotTableFields(data)
		if err != nil {
			return opt, err
		}
		opt.Data = fields
	} else {
		return opt, errors.New("data argument must be array")
	}

	filter := args.Get("filter")
	if TypeCheckOpt(filter, "array") {
		fields, err := getPivotTableFields(filter)
		if err != nil {
			return opt, err
		}
		opt.Filter = fields
	} else {
		return opt, errors.New("filter argument must be array")
	}

	rowGrandTotals := args.Get("rowGrandTotals")
	if TypeCheckOpt(rowGrandTotals, "boolean") {
		opt.RowGrandTotals = rowGrandTotals.Bool()
	} else {
		return opt, errors.New("rowGrandTotals argument must be boolean")
	}

	colGrandTotals := args.Get("colGrandTotals")
	if TypeCheckOpt(colGrandTotals, "boolean") {
		opt.ColGrandTotals = colGrandTotals.Bool()
	} else {
		return opt, errors.New("colGrandTotals argument must be boolean")
	}

	showDrill := args.Get("colGrandTotals")
	if TypeCheckOpt(showDrill, "boolean") {
		opt.ShowDrill = showDrill.Bool()
	} else {
		return opt, errors.New("showDrill argument must be boolean")
	}

	useAutoFormatting := args.Get("useAutoFormatting")
	if TypeCheckOpt(useAutoFormatting, "boolean") {
		opt.UseAutoFormatting = useAutoFormatting.Bool()
	} else {
		return opt, errors.New("useAutoFormatting argument must be boolean")
	}

	pageOverThenDown := args.Get("pageOverThenDown")
	if TypeCheckOpt(pageOverThenDown, "boolean") {
		opt.PageOverThenDown = pageOverThenDown.Bool()
	} else {
		return opt, errors.New("pageOverThenDown argument must be boolean")
	}

	mergeItem := args.Get("mergeItem")
	if TypeCheckOpt(mergeItem, "boolean") {
		opt.MergeItem = mergeItem.Bool()
	} else {
		return opt, errors.New("mergeItem argument must be boolean")
	}

	compactData := args.Get("compactData")
	if TypeCheckOpt(compactData, "boolean") {
		opt.CompactData = compactData.Bool()
	} else {
		return opt, errors.New("compactData argument must be boolean")
	}

	showError := args.Get("showError")
	if TypeCheckOpt(showError, "boolean") {
		opt.ShowError = showError.Bool()
	} else {
		return opt, errors.New("showError argument must be boolean")
	}

	showRowHeaders := args.Get("showRowHeaders")
	if TypeCheckOpt(showRowHeaders, "boolean") {
		opt.ShowRowHeaders = showRowHeaders.Bool()
	} else {
		return opt, errors.New("showRowHeaders argument must be boolean")
	}

	showColHeaders := args.Get("showColHeaders")
	if TypeCheckOpt(showColHeaders, "boolean") {
		opt.ShowColHeaders = showColHeaders.Bool()
	} else {
		return opt, errors.New("showColHeaders argument must be boolean")
	}

	showRowStripes := args.Get("showRowStripes")
	if TypeCheckOpt(showRowStripes, "boolean") {
		opt.ShowRowStripes = showRowStripes.Bool()
	} else {
		return opt, errors.New("showRowStripes argument must be boolean")
	}

	showColStripes := args.Get("showColStripes")
	if TypeCheckOpt(showColStripes, "boolean") {
		opt.ShowColStripes = showColStripes.Bool()
	} else {
		return opt, errors.New("showColStripes argument must be boolean")
	}

	showLastColumn := args.Get("showLastColumn")
	if TypeCheckOpt(showLastColumn, "boolean") {
		opt.ShowLastColumn = showLastColumn.Bool()
	} else {
		return opt, errors.New("showLastColumn argument must be boolean")
	}

	pivotTableStyleName := args.Get("pivotTableStyleName")
	if TypeCheckOpt(pivotTableStyleName, "string") {
		opt.PivotTableStyleName = pivotTableStyleName.String()
	} else {
		return opt, errors.New("pivotTableStyleName argument must be string")
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

func getPivotTableFields(args js.Value) ([]excelize.PivotTableField, error) {
	var pivotTableFields []excelize.PivotTableField

	for i := 0; i < args.Length(); i++ {
		arg := args.Index(i)
		pivotTableField, err := getPivotTableField(arg)
		if err != nil {
			return pivotTableFields, err
		}
		pivotTableFields = append(pivotTableFields, pivotTableField)	
	}
	return pivotTableFields, nil
}
