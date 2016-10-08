package conf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jingweno/conf"
)

var avg *conf.Conf
var homePath string
var configPath string

func init() {
	var err error
	if homePath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	_ = workPath
	configPath = filepath.Join(workPath, "conf", "conf.json")
	load()
}

func load() {
	loader := conf.NewLoader()
	loader.Env().Argv()
	loader.File(configPath)
	c, err := loader.Load()
	if err != nil {
		log.Panic("load conf.json error.", err)
	}
	avg = c
}

//Get 获取配置
func Get(k string) interface{} {
	return avg.Get(k)
}

//Set 添加配置
func Set(k string, v interface{}) {
	avg.Set(k, v)

}

//Bool 获取
func Bool(k string) bool {
	return avg.Bool(k)
}

//Int 获取
func Int(k string) int {
	return avg.Int(k)
}

//Merge 合并
func Merge(m map[string]interface{}) {
	avg.Merge(m)
}
