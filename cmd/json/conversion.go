package json

import (
	"errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

const (
	bufSize = 1024 * 1024
	initStructName = "Object"
)

var typeEmptyErr = errors.New("the go type is empty")

type ConversionFiles struct {
	JSONSrcFiles     []string
	GoSrcFiles       []string
	JSONSuccessFiles []string
	GoSuccessFiles   []string
}

func (f *ConversionFiles) Convert() error {
	if len(f.JSONSrcFiles) > 0 {
		if err := f.createGoFile(); err != nil {
			return err
		}
	}

	if len(f.GoSrcFiles) > 0 {
		return f.createJSONFile()
	}

	return nil
}

func (f *ConversionFiles) createGoFile() error {
	for _, jsonFile := range f.JSONSrcFiles {
		if err := oneGoFile(jsonFile); err != nil {
			return err
		}
	}

	return nil
}

/*
	1. 反序列化为 map
	2. 遍历 map
	3. 根据 value 的类型，生成 struct 类型
		数字有小数部分的，一律是 float64 类型
*/
func oneGoFile(jsonFile string) error {
	data, _ := ioutil.ReadFile(jsonFile)
	

	goFile, err := os.Create(strings.ReplaceAll(jsonFile, "json", "go"))
	if err != nil {
		return err
	}
	defer goFile.Close()

	buf := bytes.NewBuffer(make([]byte, 0, bufSize))
	if err := oneStruct(buf, data, initStructName); err != nil {
		return err
	}
	goFile.Write(buf.Bytes())

	return nil
}

// 递归
func oneStruct(buf *bytes.Buffer, data []byte, structName string) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	if structName == initStructName {
		buf.WriteString(fmt.Sprintf("type %s struct {", structName))
	} else {
		buf.WriteString(fmt.Sprintf("%s struct {", structName))
	}
	
	for jsonKey, jsonValue := range m {
		goType := goTypeString(jsonKey, jsonValue)
		if goType == "" {
			return typeEmptyErr
		}
		if goType == "[]byte" {
			if err := oneStruct(buf, jsonValue.([]byte), jsonKey); err != nil {
				return err
			}
			continue
		}
		goMember := underline2hump(jsonKey)
		buf.WriteString(fmt.Sprintf("%s %s `json:\"name:%s\"`\n", goMember, goType, jsonKey))
	}
	buf.WriteString("}")
	return nil
}

//
func goTypeString(jsonKey string, jsonValue interface{}) string {
	switch v := jsonValue.(type) {
	case float64:
		if haveIDString(jsonKey) || v-math.Ceil(v) == 0 {
			return "uint32"
		}
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case []byte:
		return "[]byte"
	}
	return ""
}

// like user_name to UserName
// id will be ID
// json will be JSON
func underline2hump(s string) string {
	s = strings.ToLower(s)
	
	if strings.Contains(s, "json") {
		s = strings.ReplaceAll(s, "json", "JSON")
	}

	if strings.Contains(s, "id") {
		s = strings.ReplaceAll(s, "id", "ID")
	}

	strs := strings.Split(s, "_")
	var builder strings.Builder
	for _, str := range strs {
		data := []byte(str)
		if data[0] > 90 {
			data[0] -= 32
		}
		builder.Write(data)
	}

	return builder.String()
}

func haveIDString(s string) bool {
	s = strings.ToLower(s)
	return strings.Contains(s, "id")
}

func (f *ConversionFiles) createJSONFile() error {
	return nil
}
