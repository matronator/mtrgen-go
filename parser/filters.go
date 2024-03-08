package parser

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Filters is a struct that contains filter functions
type Filters struct{}

// ENCODING is the encoding used by filter functions
const ENCODING = "UTF-8"

// GLOBAL_FILTERS is a slice containing global filter names
var GLOBAL_FILTERS = []string{"upper", "lower", "upperFirst", "lowerFirst", "first", "last", "camelCase", "snakeCase", "kebabCase", "pascalCase", "titleCase", "length", "reverse", "random", "truncate"}

func (f *Filters) ApplyFilter(filter string, str string, args ...any) string {
	switch filter {
	case "upper":
		return f.Upper(str)
	case "lower":
		return f.Lower(str)
	case "upperFirst":
		return f.UpperFirst(str)
	case "lowerFirst":
		return f.LowerFirst(str)
	case "first":
		return f.First(str, args[0].(int))
	case "last":
		return f.Last(str, args[0].(int))
	case "camelCase":
		return f.CamelCase(str)
	case "snakeCase":
		return f.SnakeCase(str)
	case "kebabCase":
		return f.KebabCase(str)
	case "pascalCase":
		return f.PascalCase(str)
	case "titleCase":
		return f.TitleCase(str)
	case "length":
		return f.Length(str)
	case "reverse":
		return f.Reverse(str)
	case "random":
		return f.Random(str)
	case "truncate":
		return f.Truncate(str, args[0].(int))
	default:
		return str
	}
}

// Upper converts a string to uppercase using the specified encoding
func (f *Filters) Upper(str string) string {
	return strings.ToUpper(str)
}

// Lower converts a string to lowercase using the specified encoding
func (f *Filters) Lower(str string) string {
	return strings.ToLower(str)
}

// UpperFirst converts the first character of a string to uppercase
func (f *Filters) UpperFirst(str string) string {
	if str == "" {
		return str
	}
	sep := " "
	ss := strings.SplitN(str, sep, 2)
	r := []rune(ss[0])
	r[0] = unicode.ToUpper(r[0])
	str = string(r)
	if len(ss) > 1 {
		str += sep + ss[1]
	}

	return str
}

// LowerFirst converts the first character of a string to lowercase
func (f *Filters) LowerFirst(str string) string {
	if str == "" {
		return str
	}
	sep := " "
	ss := strings.SplitN(str, sep, 2)
	r := []rune(ss[0])
	r[0] = unicode.ToLower(r[0])
	str = string(r)
	if len(ss) > 1 {
		str += sep + ss[1]
	}

	return str
}

// First returns the first n characters of a string
func (f *Filters) First(str string, n int) string {
	if n > len(str) {
		n = len(str)
	}
	return str[:n]
}

// Last returns the last n characters of a string
func (f *Filters) Last(str string, n int) string {
	if n > len(str) {
		n = len(str)
	}
	return str[len(str)-n:]
}

// CamelCase converts a string to camel case
func (f *Filters) CamelCase(str string) string {
	words := strings.Fields(str)
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = cases.Title(language.Und, cases.NoLower).String(words[i])
		}
	}
	return strings.Join(words, "")
}

// SnakeCase converts a string to snake case
func (f *Filters) SnakeCase(str string) string {
	return strings.Join(strings.FieldsFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}), "_")
}

// KebabCase converts a string to kebab case
func (f *Filters) KebabCase(str string) string {
	return strings.Join(strings.FieldsFunc(str, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}), "-")
}

// PascalCase converts a string to pascal case
func (f *Filters) PascalCase(str string) string {
	words := strings.Fields(str)
	for i := range words {
		words[i] = cases.Title(language.Und, cases.NoLower).String(words[i])
	}
	return strings.Join(words, "")
}

// TitleCase converts a string to title case
func (f *Filters) TitleCase(str string) string {
	return cases.Title(language.Und, cases.NoLower).String(str)
}

// Length returns the length of a string
func (f *Filters) Length(str string) string {
	return strconv.Itoa(len(str))
}

// Reverse reverses the characters in a string
func (f *Filters) Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Random returns a random character from a string
func (f *Filters) Random(str string) string {
	if str == "" {
		return str
	}
	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(len(str))
	return string(str[randomIndex])
}

// Truncate truncates a string to the specified length
func (f *Filters) Truncate(str string, length int) string {
	if length >= len(str) {
		return str
	}
	return str[:length]
}
