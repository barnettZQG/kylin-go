package kylin

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type DataBase struct {
	ProjectName string
}

//QueryAll 查询主方法
func (d *DataBase) QueryAll(tableName string, where interface{}, offset, limit int) ([]byte, error) {
	tableName = strings.ToUpper(tableName)
	reflect.TypeOf(where)
	s := reflect.ValueOf(where).Elem()
	st := reflect.TypeOf(where).Elem()
	queryWhere := "Select * from " + tableName + " where 1=1 "
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ft := st.Field(i)
		if f.Type().String() == "int" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(int) == 0 {
				return nil, fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if f.Interface().(int) != 0 {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "=" + strconv.Itoa(f.Interface().(int))
			}
		}
		if f.Type().String() == "string" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(string) == "" {
				return nil, fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if f.Interface().(string) != "" {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "='" + f.Interface().(string) + "'"
			}
		}
		if f.Type().String() == "float64" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(float64) == 0.0 {
				return nil, fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if f.Interface().(float64) != 0.0 {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "=" + strconv.FormatFloat(f.Interface().(float64), 'g', 2, 64)
			}
		}
		if f.Type().String() == "time.Time" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(time.Time).IsZero() {
				return nil, fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if !f.Interface().(time.Time).IsZero() {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + ">='" + f.Interface().(time.Time).Format("2006-01-01") + "'"
			}
		}
	}
	fmt.Println(queryWhere)
	query := &Query{
		SQL:     queryWhere,
		Limit:   limit,
		Offset:  offset,
		Project: d.ProjectName,
	}
	code, body, err := QueryKylin(query)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("kylin server return error:" + string(body))
	}
	return body, nil
}

//QueryOne 查询一个
func (d *DataBase) QueryOne(where interface{}) interface{} {
	return nil
}
