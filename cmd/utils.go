package main

import (
	"errors"
	"syscall/js"
)

// expectType expect js type string, number, boolean, null, undefined, object, array, function
func TypeCheck(jsValue js.Value, expectType string) (bool, error) {
	if expectType == "string" {
		return jsValue.Type() == js.TypeString, nil
	} else if expectType == "number" {
		return jsValue.Type() == js.TypeNumber, nil
	} else if expectType == "boolean" {
		return jsValue.Type() == js.TypeBoolean, nil
	} else if expectType == "null" {
		return jsValue.Type() == js.TypeNull, nil
	} else if expectType == "undefined" {
		return jsValue.Type() == js.TypeUndefined, nil
	} else if expectType == "array" {
		return jsValue.Type() == js.TypeObject && jsValue.Length() >= 0, nil
	} else if expectType == "object" {
		return jsValue.Type() == js.TypeObject, nil
	} else if expectType == "function" {
		return jsValue.Type() == js.TypeFunction, nil
	}
	return false, errors.New("No such expect error")
}

func TypeCheckOpt(jsValue js.Value, expectType string) (bool, error) {
	if jsValue.Type() == js.TypeUndefined {
		return true, nil
	}
	return TypeCheck(jsValue, expectType)
}