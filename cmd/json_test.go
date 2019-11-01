package cmd

import (
	"strings"
	"testing"
)

type MockGetAllFiles func() []string

func testcheckFileName(mock MockGetAllFiles) bool {
	files := mock()
	if files == nil || len(files) <= 0 {
		return false
	}

	files = getPrefix(files)

	set := map[string]struct{}{}
	for _, f := range files {
		set[f] = struct{}{}
	}

	return len(set) == len(files)
}

func TestCheckFileName(t *testing.T) {
	testCases := []struct {
		f    MockGetAllFiles
		want bool
	}{
		{
			func() []string {
				return nil
			},
			false,
		},
		{
			func() []string {
				return []string{"a", "a"}
			},
			false,
		},
		{
			func() []string {
				return []string{"a", "b"}
			},
			true,
		},
		{
			func() []string {
				return []string{"a.go", "a.json"}
			},
			false,
		},
		{
			func() []string {
				return []string{"a.json", "b.json"}
			},
			true,
		},
	}

	for _, testCase := range testCases {
		if testcheckFileName(testCase.f) != testCase.want {
			t.Error("test fail")
		}
	}
}

func TestReadDir(t *testing.T) {
	t.Logf("%v\n", getAllFile())
}

func TestAA(t *testing.T) {
	t.Logf("%v\n", strings.TrimSuffix("aaa.json", ".go"))
}

func TestBB(t *testing.T) {
	t.Logf("%v\n", strings.ReplaceAll("a.json", "json", "go"))
}
