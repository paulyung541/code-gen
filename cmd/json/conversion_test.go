package json

import (
	"bytes"
	"testing"
)

func TestBuf(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	buf.WriteString("type Object struct {\n")
	buf.WriteString("	name int `json:\"name\"`\n")
	buf.WriteString("}")
	t.Log(buf.String())
}

func TestOneStruct(t *testing.T) {
	jsonData1 := []byte(`{
    "id": 53,
    "url": "localhost",
	"yes": false
}
`)
	buf := bytes.NewBuffer(make([]byte, bufSize))
	oneStruct(buf, jsonData1, initStructName)
	t.Log(buf.String())

	jsonData2 := []byte(`{
    "id": 100,
    "data": {
        "name": "xiaoming"
    }
}`)

	buf.Reset()
	oneStruct(buf, jsonData2, initStructName)
	t.Log(buf.String())

}

func TestUnderline2hump(t *testing.T) {
	testCase := []struct {
		str  string
		want string
	}{
		{
			"user_name",
			"userName",
		},
		{
			"good_things_i_think",
			"goodThingsIThink",
		},
		{
			"name_id",
			"nameID",
		},
		{
			"test_json_file",
			"testJSONFile",
		},
	}

	for _, c := range testCase {
		if fact := underline2hump(c.str); fact != c.want {
			t.Errorf("fail: c.str = %s, want = %s, fact = %s", c.str, c.want, fact)
		}
	}
}
