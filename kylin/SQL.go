package kylin

import (
	"fmt"
	"kylin-orm/models"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type sql struct {
	sql          string
	tableName    string
	limit        int
	offset       int
	wheres       []*where
	selects      []selects
	groups       []group
	orders       []*order
	isError      bool
	errorMessage string
}
type where struct {
	Key    string
	Value  interface{}
	Symbol string
}
type selects string
type group string
type order struct {
	fun  string
	name string
	sort string
}

func SQL(tableName string) *sql {
	if tableName == "" {
		return nil
	}
	return &sql{
		tableName: tableName,
	}
}
func (s *sql) error(message string) {
	s.isError = true
	s.errorMessage = message
}

func (s *sql) Select(se ...selects) *sql {
	s.selects = append(s.selects, se...)
	return s
}
func (s *sql) Fromsql(sql string) *sql {
	s.sql = sql
	return s
}

func (s *sql) Where(key, symbol string, value interface{}) *sql {
	s.wheres = append(s.wheres, &where{
		Key:    key,
		Value:  value,
		Symbol: symbol,
	})
	return s
}

func (s *sql) WhereAll(w interface{}) *sql {
	v := reflect.ValueOf(w).Elem()
	t := reflect.TypeOf(w).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		//忽略字段
		if ft.Tag.Get("kylin") == "nowhere" {
			continue
		}
		if f.Type().String() == "int" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(int) == 0 {
				s.error(ft.Name + "字段是必须的字段")
			} else if f.Interface().(int) != 0 {
				s.wheres = append(s.wheres, &where{
					Key:    s.tableName + "." + strings.ToUpper(ft.Tag.Get("json")),
					Value:  f.Interface().(int),
					Symbol: "=",
				})
			}
		}
		if f.Type().String() == "string" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(string) == "" {
				s.error(ft.Name + "字段是必须的字段")
			} else if f.Interface().(string) != "" {
				s.wheres = append(s.wheres, &where{
					Key:    s.tableName + "." + strings.ToUpper(ft.Tag.Get("json")),
					Value:  f.Interface().(string),
					Symbol: "=",
				})
			}
		}
		if f.Type().String() == "float64" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(float64) == 0.0 {
				s.error(ft.Name + "字段是必须的字段")
			} else if f.Interface().(float64) != 0.0 {
				s.wheres = append(s.wheres, &where{
					Key:    s.tableName + "." + strings.ToUpper(ft.Tag.Get("json")),
					Value:  f.Interface().(float64),
					Symbol: "=",
				})
			}
		}
		if f.Type().String() == "models.KylinTime" {
			if ft.Tag.Get("kylin") == "necessary" && f.Interface().(models.KylinTime).IsZero() {
				s.error(ft.Name + "字段是必须的字段")
			} else if !f.Interface().(models.KylinTime).IsZero() {
				s.wheres = append(s.wheres, &where{
					Key:    s.tableName + "." + strings.ToUpper(ft.Tag.Get("json")),
					Value:  f.Interface().(models.KylinTime).StartTime,
					Symbol: ">=",
				})
				s.wheres = append(s.wheres, &where{
					Key:    s.tableName + "." + strings.ToUpper(ft.Tag.Get("json")),
					Value:  f.Interface().(models.KylinTime).EndTime,
					Symbol: "<",
				})
			}
		}
	}
	return s
}

func (s *sql) Group(g ...group) *sql {
	s.groups = append(s.groups, g...)
	return s
}
func (s *sql) Order(name, sortType, fun string) *sql {
	if strings.ToUpper(sortType) != "DESC" && strings.ToUpper(sortType) != "ASC" {
		s.error("sortType  error, could to be :DESC and ASC ")
	}
	s.orders = append(s.orders, &order{
		fun:  fun,
		name: name,
		sort: sortType,
	})
	return s
}
func (s *sql) Build() (*sql, error) {
	if s.isError {
		return nil, fmt.Errorf(s.errorMessage)
	}
	q := "Select "
	if s.selects == nil || len(s.selects) < 1 {
		q += " * "
	} else {
		for _, ss := range s.selects {
			q += " " + string(ss) + ","
		}
		if strings.HasSuffix(q, ",") {
			q = q[0 : len(q)-1]
		}
	}
	q += " From " + strings.ToUpper(s.tableName) + " "
	if s.wheres != nil && len(s.wheres) > 0 {
		q += " Where 1=1 "
		for _, w := range s.wheres {
			name := reflect.TypeOf(w.Value).Name()
			//fmt.Println(name)
			if name == "string" {
				q += " and " + w.Key + w.Symbol + "'" + w.Value.(string) + "'"
			} else if name == "int" {
				q += " and " + w.Key + w.Symbol + strconv.Itoa(w.Value.(int))
			} else if name == "float64" {
				q += " and " + w.Key + w.Symbol + strconv.FormatFloat(w.Value.(float64), 'g', 2, 64)
			} else if name == "Time" {
				q += " and " + w.Key + w.Symbol + "'" + w.Value.(time.Time).Format("2006-01-02") + "'"
			}
		}
	}
	if s.groups != nil && len(s.groups) > 0 {
		q += " Group By ("
		for _, sg := range s.groups {
			q += s.tableName + "." + strings.ToUpper(string(sg)) + ","
		}
		if strings.HasSuffix(q, ",") {
			q = q[0 : len(q)-1]
		}
		q += ") "
	}
	if s.orders != nil && len(s.orders) > 0 {
		q += " Order By "
		for _, so := range s.orders {
			if so.fun != "" {
				q += so.fun + "(" + s.tableName + "." + strings.ToUpper(so.name) + ") " + strings.ToUpper(so.sort) + ","
			} else {
				q += s.tableName + "." + strings.ToUpper(so.name) + " " + strings.ToUpper(so.sort) + ","
			}

		}
		if strings.HasSuffix(q, ",") {
			q = q[0 : len(q)-1]
		}
	}
	s.sql = q
	return s, nil
}
func (s *sql) String() string {
	if s.sql == "" {
		s.Build()
	}
	return s.sql
}
func (s *sql) Limit(limit int) *sql {
	s.limit = limit
	return s
}
func (s *sql) Offset(offset int) *sql {
	s.offset = offset
	return s
}

func (s *sql) GetLimit() int {
	if s.limit == 0 {
		return 10
	}
	return s.limit
}
func (s *sql) GetOffset() int {
	return s.offset
}
