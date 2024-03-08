// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package template

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/oriser/regroup"
	"mtrgen/parser"
)

const HeaderPattern = `(?sm)^\S+ --- MTRGEN ---.(?P<fields>.+)\s\S+ --- MTRGEN ---`

type HeaderMatch struct {
	Fields string `regroup:"fields"`
}

func TransformFile(path string, arguments parser.Argument) *GenericFileObject {
	parsed := parser.ParseFile(path, arguments)
	header := GetTemplateHeader(parsed)
	parsed = RemoveHeader(parsed)

	return NewFileObject(header.Name, filepath.Clean(header.Path), parsed)
}

func WriteFiles(files []*GenericFileObject) {
	for _, file := range files {
		writeFile(file)
	}
}

func GetTemplateHeader(content string) *Header {
	var pattern = regroup.MustCompile(HeaderPattern)

	match := &HeaderMatch{}

	err := pattern.MatchToTarget(content, match)

	if err != nil {
		panic(err)
	}

	lines := regexp.MustCompile("\r\n|\n").Split(content, -1)

	info := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		kv := strings.Split(line, ":")
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		info[k] = v
	}

	if info["name"] != "" || info["filename"] != "" || info["path"] != "" {
		panic("Template header is missing some required properties (name, filename, path).")
	}

	return FromMap(info)
}

func RemoveHeader(content string) string {
	return strings.TrimSpace(regexp.MustCompile(HeaderPattern).ReplaceAllString(content, ""))
}

func writeFile(gfo *GenericFileObject) {
	_ = os.MkdirAll(filepath.Dir(gfo.Path), os.ModePerm)

	_ = os.WriteFile(gfo.Filename, []byte(gfo.Contents), os.ModePerm)
}
