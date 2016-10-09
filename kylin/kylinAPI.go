package kylin

import (
	"encoding/base64"
	"fmt"
	"kylin-orm/conf"

	"strconv"

	"github.com/valyala/fasthttp"
)

var (
	url, user, password string
)

func init() {
	url = conf.Get("kylin_url").(string)
	if url == "" {
		panic("kylin url is not empty")
	}
	user = conf.Get("kylin_user").(string)
	if user == "" {
		panic("kylin user is not empty")
	}
	password = conf.Get("kylin_pass").(string)
	if password == "" {
		panic("kylin password is not empty")
	}
	fmt.Println("Kylin_rul:", url, "Kylin_user:", user, "Kylin_pass:", password)
}

func getRequest() *fasthttp.Request {
	request := fasthttp.AcquireRequest()
	request.Header.Set("Authorization", "Basic "+basicAuth(user, password))
	return request
}

func do(request *fasthttp.Request) (code int, body []byte, err error) {
	response := fasthttp.AcquireResponse()
	err = fasthttp.Do(request, response)
	if err == nil {
		code = response.StatusCode()
		body = response.Body()
	}
	return
}
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

//Login 登陆
func Login() (code int, body []byte, err error) {
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/user/authentication")
	request.Header.SetMethod("POST")
	return do(request)
}

//QueryKylin 查询
func QueryKylin(query *Query) (code int, body []byte, err error) {
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/query")
	request.Header.SetMethod("POST")
	request.Header.Set("Content-Type", "application/json")
	request.SetBody(query.GetBytes())
	return do(request)
}

//ListTables 列出tables
//project 必须参数
func ListTables(project string) (code int, body []byte, err error) {
	if project == "" {
		return 0, nil, fmt.Errorf("project is not empty")
	}
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/tables_and_columns")
	request.PostArgs().Add("project", project)
	request.Header.SetMethod("GET")
	return do(request)
}

//ListCubes 列出cubes
//offset limit 必须参数
//cubeName, projectName 可选参数
func ListCubes(offset, limit int, cubeName, projectName string) (code int, body []byte, err error) {
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/cubes")
	request.Header.SetMethod("GET")
	request.PostArgs().Add("offset", strconv.Itoa(offset))
	request.PostArgs().Add("limit", strconv.Itoa(limit))
	request.PostArgs().Add("cubeName", cubeName)
	request.PostArgs().Add("projectName", projectName)
	return do(request)
}

//GetCube 获取cube
//cubeName 必须参数
func GetCube(cubeName string) (code int, body []byte, err error) {
	if cubeName == "" {
		return 0, nil, fmt.Errorf("cubeName is not empty")
	}
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/cubes/" + cubeName)
	request.Header.SetMethod("GET")
	return do(request)
}

//GetCubeDesc 获取cube描述
//cubeName 必须参数
func GetCubeDesc(cubeName string) (code int, body []byte, err error) {
	if cubeName == "" {
		return 0, nil, fmt.Errorf("cubeName is not empty")
	}
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/cube_desc/" + cubeName)
	request.Header.SetMethod("GET")
	return do(request)
}

//GetModel 获取model
//modelName 必须参数
func GetModel(modelName string) (code int, body []byte, err error) {
	if modelName == "" {
		return 0, nil, fmt.Errorf("modelName is not empty")
	}
	request := getRequest()
	request.SetRequestURI(url + "/kylin/api/model/" + modelName)
	request.Header.SetMethod("GET")
	return do(request)
}

// //CreateCube 创建cube
// func CreateCube(cubeName string) (code int, body []byte, err error) {
// 	request := getRequest()
// 	request.SetRequestURI(url + "/kylin/api/cubes/" + cubeName + "/rebuild")
// 	request.Header.SetMethod("PUT")
// 	return do(request)
// }
