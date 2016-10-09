package kylin

import (
	"encoding/json"
	"fmt"
)

//Query 查询条件封装
type Query struct {
	SQL           string `json:"sql"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	AcceptPartial bool   `json:"acceptPartial"`
	Project       string `json:"project"`
}

//GetBytes 对象转json
func (query *Query) GetBytes() (body []byte) {
	var err error
	body, err = json.Marshal(query)
	if err != nil {
		fmt.Println("query to json error.", err)
		return nil
	}
	return
}
