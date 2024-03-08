package parser

import (
	"bytes"
	"os"
	"strconv"
	"strings"

	"github.com/oriser/regroup"
	"mtrgen/utils"
)

const VariablePattern = `(?m)(?P<all><%\s?(?P<name>[a-zA-Z0-9_]+)(?P<default>=.+?)?\|?(?P<filterWithArgs>(?P<filter>[a-zA-Z0-9_]+?)(?:\:(?:(?:\\?\'|\\?")?.?(?:\\?\'|\\?")?,?)+?)*?)?\s?%>)`

const LiterallyNull = "__:-LITERALLY_NULL-:__"

type Argument map[string]interface{}

type Matches struct {
	EntireMatch    string `regroup:"all"`
	Name           string `regroup:"name"`
	DefaultValue   string `regroup:"default"`
	FilterWithArgs string `regroup:"filterWithArgs"`
	Filter         string `regroup:"filter"`
}

func ParseString(str string, arguments Argument) string {
	var pattern = regroup.MustCompile(VariablePattern)

	matches := &Matches{}
	rets, err := pattern.MatchAllToTarget(str, -1, matches)

	if err != nil {
		panic(err)
	}

	args := make(Argument, len(rets))

	for _, match := range rets {
		name := match.(*Matches).Name
		if arguments[name] != nil {
			args[name] = arguments[name]
		} else {
			defaultVal := getDefaultValue(match.(*Matches).DefaultValue)
			if defaultVal != LiterallyNull {
				args[name] = defaultVal
			} else {
				args[name] = nil
			}
		}
	}

	retsMap := make(map[string]Matches, len(rets))
	for _, match := range rets {
		retsMap[match.(*Matches).Name] = *match.(*Matches)
	}

	modified := applyFilters(retsMap, args)

	argArray := make([]string, len(modified))
	for _, value := range modified {
		argArray = append(argArray, value.(string))
	}

	retsArray := make([]string, len(rets))
	for _, value := range rets {
		retsArray = append(retsArray, value.(*Matches).EntireMatch)
	}

	return strings.NewReplacer(utils.Zip(retsArray, argArray)...).Replace(str)
}

func ParseFile(filename string, arguments Argument) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	file = bytes.TrimSpace(file)

	return ParseString(string(file), arguments)
}

func getDefaultValue(str string) any {
	defaultValue := strings.Trim(str, "='\"`")

	if defaultValue == "" {
		return LiterallyNull
	}

	if defaultValue == "false" {
		return false
	} else if defaultValue == "true" {
		return true
	} else if defaultValue == "null" {
		return nil
	} else if utils.IsNum(defaultValue) {
		i := strings.Index(defaultValue, ".")
		if i != -1 {
			val, _ := strconv.ParseFloat(defaultValue, 64)

			return val
		}

		val, _ := strconv.ParseInt(defaultValue, 10, 64)

		return val
	}

	return defaultValue
}

func applyFilters(matches map[string]Matches, arguments Argument) Argument {
	modified := arguments

	var args []any

	for key, arg := range arguments {
		if matches[key].Filter != "" {
			filter := matches[key].Filter
			if matches[key].FilterWithArgs != filter {
				filterArgumentsString := strings.SplitN(matches[key].FilterWithArgs, ":", 2)[1]
				filterArguments := strings.Split(filterArgumentsString, ",")
				args = mapArgs(filterArguments)
			} else {
				args = []any{arg}
			}

			filters := Filters{}
			modified[key] = filters.ApplyFilter(filter, arguments[key].(string), args...)

			// filterFunc := GetFilterFunction(filter)
			//
			// if !filterFunc.IsZero() {
			// 	// Convert args to []reflect.Value
			// 	reflectArgs := make([]reflect.Value, len(args))
			// 	for i, arg := range args {
			// 		reflectArgs[i] = reflect.ValueOf(arg)
			// 	}
			//
			// 	modified[key] = filterFunc.Call(reflectArgs)[0].Interface()
			// } else {
			// 	panic("Filter function '" + filter + "' does not exist.")
			// }
		}
	}

	return modified
}

func mapArgs(args []string) []interface{} {
	mappedArgs := make([]interface{}, len(args))

	for i, item := range args {
		item = strings.TrimSpace(item)

		switch {
		case utils.IsNum(item) && !strings.Contains(item, "."):
			mappedArgs[i], _ = strconv.Atoi(item)
		case utils.IsNum(item) && strings.Contains(item, "."):
			mappedArgs[i], _ = strconv.ParseFloat(item, 32)
		case item == "true":
			mappedArgs[i] = true
		case item == "false":
			mappedArgs[i] = false
		case item == "null":
			mappedArgs[i] = nil
		default:
			mappedArgs[i] = strings.Trim(item, "\"'")
		}
	}

	return mappedArgs
}
