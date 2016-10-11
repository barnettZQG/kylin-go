package kylin

import (
	"encoding/json"
	"fmt"
	"kylin-orm/models"
	"reflect"
	"strconv"
	"strings"
)

type DataBase struct {
	ProjectName string
}

//QueryAll 查询主方法
func (d *DataBase) QueryAll(tableName string, where interface{}, offset, limit int) (*QueryResult, error) {
	return d.QueryPart(tableName, nil, where, offset, limit)
}

//QueryOne 查询一个
func (d *DataBase) QueryOne(tableName string, where interface{}) (*QueryResult, error) {
	return d.QueryPart(tableName, nil, where, 0, 1)
}

//QueryPart 查询部分字段
func (d *DataBase) QueryPart(tableName string, fields []string, where interface{}, offset, limit int) (*QueryResult, error) {
	queryWhere, err := d.getSQL(fields, where, tableName)
	if err != nil {
		return nil, err
	}
	//fmt.Println(queryWhere)
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
	return d.handleBody(body)
}

//QueryBySQL 通过sql查询
func (d *DataBase) QueryBySQL(sql *sql) (*QueryResult, error) {

	//fmt.Println(queryWhere)
	query := &Query{
		SQL:     sql.String(),
		Limit:   sql.GetLimit(),
		Offset:  sql.GetOffset(),
		Project: d.ProjectName,
	}
	code, body, err := QueryKylin(query)
	if err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("kylin server return error:" + string(body))
	}
	return d.handleBody(body)
}

//GetSQL 构建查询sql
func (d *DataBase) getSQL(fields []string, where interface{}, tableName string) (string, error) {
	tableName = strings.ToUpper(tableName)
	queryWhere := "Select "
	if fields == nil {
		queryWhere += " * "
	} else {
		for _, field := range fields {
			queryWhere += strings.ToUpper(field) + ","
		}
		queryWhere = queryWhere[0 : len(queryWhere)-1]
	}
	s := reflect.ValueOf(where).Elem()
	st := reflect.TypeOf(where).Elem()
	queryWhere += " from " + tableName + " where 1=1 "
	for i := 0; i < s.NumField(); i++ {

		f := s.Field(i)
		ft := st.Field(i)
		//忽略字段
		if ft.Tag.Get("kylin") == "nowhere" {
			continue
		}
		if f.Type().String() == "int" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(int) == 0 {
				return "", fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if f.Interface().(int) != 0 {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "=" + strconv.Itoa(f.Interface().(int))
			}
		}
		if f.Type().String() == "string" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(string) == "" {
				return "", fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if f.Interface().(string) != "" {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "='" + f.Interface().(string) + "'"
			}
		}
		if f.Type().String() == "float64" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(float64) == 0.0 {
				return "", fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if f.Interface().(float64) != 0.0 {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "=" + strconv.FormatFloat(f.Interface().(float64), 'g', 2, 64)
			}
		}
		if f.Type().String() == "models.KylinTime" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(models.KylinTime).IsZero() {
				return "", fmt.Errorf(ft.Name + "字段是必须的字段")
			} else if !f.Interface().(models.KylinTime).IsZero() {
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + ">='" + f.Interface().(models.KylinTime).StartTime.Format("2006-01-02") + "'"
				queryWhere = queryWhere + " and " + tableName + "." + strings.ToUpper(ft.Tag.Get("json")) + "<'" + f.Interface().(models.KylinTime).EndTime.Format("2006-01-02") + "'"
			}
		}
	}
	return queryWhere, nil
}

func (d *DataBase) handleBody(body []byte) (*QueryResult, error) {
	if body == nil {
		return nil, nil
	}
	re := &QueryResult{}
	err := json.Unmarshal(body, re)
	if err != nil {
		return nil, err
	}
	if re.IsException {
		return nil, fmt.Errorf(re.ExceptionMessage)
	}
	return re, nil
}
