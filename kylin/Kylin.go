package kylin

import "kylin-orm/conf"

//Kylin kylin-orm interface
type Kylin interface {
	QueryAll(tableName string, where interface{}, offset, limit int) (*QueryOut, error)
	QueryOne(tableName string, where interface{}) (*QueryOut, error)
	QueryPart(tableName string, fields []string, where interface{}, offset, limit int) (*QueryOut, error)
	QueryBySQL(sql *sql) (*QueryOut, error)
}

//kylinBase 查询入口
type kylinBase struct {
	DataBase *DataBase
}

var base *kylinBase

//GetDefaultKylinBase 获取默认base
func GetDefaultKylinBase() *kylinBase {
	if base == nil {
		base = &kylinBase{
			DataBase: &DataBase{
				ProjectName: conf.String("kylin_project"),
			},
		}
	}
	return base
}

//GetKylinBase 获取指定base
func GetKylinBase(project string) *kylinBase {
	if project == "" {
		return nil
	}
	b := &kylinBase{
		DataBase: &DataBase{
			ProjectName: project,
		},
	}
	return b
}

//QueryAll 查询全部
func (k *kylinBase) QueryAll(tableName string, where interface{}, offset, limit int) (q *QueryOut, err error) {
	re, err := k.DataBase.QueryAll(tableName, where, offset, limit)
	if err != nil {
		return nil, err
	}
	var metas []string
	for _, m := range re.ColumnMetas {
		metas = append(metas, m.Name)
	}
	q = &QueryOut{
		ColumnMetas: metas,
		Result:      re.Result,
	}
	// q.ColumnMetas = metas
	// q.Result = re.Result
	return
}

//QueryPart 查询全部
func (k *kylinBase) QueryPart(tableName string, fields []string, where interface{}, offset, limit int) (q *QueryOut, err error) {
	re, err := k.DataBase.QueryPart(tableName, fields, where, offset, limit)
	if err != nil {
		return nil, err
	}
	var metas []string
	for _, m := range re.ColumnMetas {
		metas = append(metas, m.Name)
	}
	q = &QueryOut{
		ColumnMetas: metas,
		Result:      re.Result,
	}
	return
}

//QueryOne 查询一个
func (k *kylinBase) QueryOne(tableName string, where interface{}) (q *QueryOut, err error) {
	re, err := k.DataBase.QueryOne(tableName, where)
	if err != nil {
		return nil, err
	}
	var metas []string
	for _, m := range re.ColumnMetas {
		metas = append(metas, m.Name)
	}
	q = &QueryOut{
		ColumnMetas: metas,
		Result:      re.Result,
	}
	return
}

//QueryBySQL 使用构建好的sql
func (k *kylinBase) QueryBySQL(sql *sql) (q *QueryOut, err error) {
	re, err := k.DataBase.QueryBySQL(sql)
	if err != nil {
		return nil, err
	}
	var metas []string
	for _, m := range re.ColumnMetas {
		metas = append(metas, m.Name)
	}
	q = &QueryOut{
		ColumnMetas: metas,
		Result:      re.Result,
	}
	return
}
