package kylin

import (
	"fmt"
	"kylin-orm/kylin"
	"testing"
)

func showResult(code int, body []byte, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(code, string(body))
}
func TestQueryKylin(t *testing.T) {
	showResult(kylin.QueryKylin(&kylin.Query{
		SQL:           "SELECT * FROM XXX",
		Offset:        0,
		Limit:         1000,
		AcceptPartial: false,
		Project:       "dsp",
	}))
}

func TestListTables(t *testing.T) {
	showResult(kylin.ListTables("DSP"))
}

func TestListCubes(t *testing.T) {
	showResult(kylin.ListCubes(0, 10, "", ""))
}

func TestGetCube(t *testing.T) {
	showResult(kylin.GetCube("xxx"))
}

func TestGetCubeDesc(t *testing.T) {
	showResult(kylin.GetCubeDesc("xxx"))
}
func TestGetModel(t *testing.T) {
	showResult(kylin.GetModel("xxx"))
}
