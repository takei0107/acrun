package lang

import (
	"encoding/json"
	"fmt"
	"strings"
)

// -ldflags "-X 'github.com/takei0107/acrun/lang.CJson=$(cat internal/lang/filetypes/c.json)'"
var (
	CJson string
)

type JsonParam struct {
	IsNeedCompile   bool              `json:"isNeedCompile"`
	DefaultFileName string            `json:"defaultFileName"`
	DefaultExe      string            `json:"defaultExe"`
	Compile         *JsonCompileParam `json:"compile"`
	Run             *JsonExeParam     `json:"run"`
}

type JsonCompileParam struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

type JsonExeParam struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func handleJsonString(j string) (*JsonParam, error) {
	r := strings.NewReader(j)
	d := json.NewDecoder(r)

	var jp *JsonParam

	err := d.Decode(&jp)
	if err != nil {
		return nil, err
	}

	return jp, nil
}

func GetJsonParam(ft FileType) (*JsonParam, error) {
	c := ftConfigMap[ft]

	if c == nil {
		return nil, fmt.Errorf("filetype=%d is invalid\n", ft)
	}

	if c.jsonString == "" {
		return nil, fmt.Errorf("json string is empty. filetype=%d\n", ft)
	}

	jp, err := handleJsonString(c.jsonString)
	if err != nil {
		return nil, err
	}

	return jp, nil
}
