package test

// import (
// 	"fmt"
// 	"kylin-orm/kylin"
// 	"testing"
// )

// func showResult(code int, body []byte, err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(code, string(body))
// }
// func TestQueryKylin(t *testing.T) {
// 	query := &kylin.Query{
// 		SQL:           "SELECT * FROM dsp_online_report",
// 		Offset:        0,
// 		Limit:         10000,
// 		AcceptPartial: false,
// 		Project:       "dsp_online_test",
// 	}
// 	fmt.Println(string(query.GetBytes()))
// 	showResult(kylin.QueryKylin(query))
// }

// func TestListTables(t *testing.T) {
// 	showResult(kylin.ListTables("dsp_online_test"))
// }

// func TestListCubes(t *testing.T) {
// 	showResult(kylin.ListCubes(0, 10, "", ""))
// }

// func TestGetCube(t *testing.T) {
// 	showResult(kylin.GetCube("xxx"))
// }

// func TestGetCubeDesc(t *testing.T) {
// 	showResult(kylin.GetCubeDesc("xxx"))
// }
// func TestGetModel(t *testing.T) {
// 	showResult(kylin.GetModel("xxx"))
// }
