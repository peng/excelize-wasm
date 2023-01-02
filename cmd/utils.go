package main

import (
	"syscall/js"
)

// expectType expect js type string, number, boolean, null, undefined, object, array, function
func TypeCheck(jsValue js.Value, expectType string) bool {
	if expectType == "string" {
		return jsValue.Type() == js.TypeString
	} else if expectType == "number" {
		return jsValue.Type() == js.TypeNumber
	} else if expectType == "boolean" {
		return jsValue.Type() == js.TypeBoolean
	} else if expectType == "null" {
		return jsValue.Type() == js.TypeNull
	} else if expectType == "undefined" {
		return jsValue.Type() == js.TypeUndefined
	} else if expectType == "array" {
		return jsValue.Type() == js.TypeObject && jsValue.Length() >= 0
	} else if expectType == "object" {
		return jsValue.Type() == js.TypeObject
	} else if expectType == "function" {
		return jsValue.Type() == js.TypeFunction
	}
	return false
}

func TypeCheckOpt(jsValue js.Value, expectType string) bool {
	if jsValue.Type() == js.TypeUndefined {
		return true
	}
	return TypeCheck(jsValue, expectType)
}