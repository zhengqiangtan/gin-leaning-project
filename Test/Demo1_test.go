package Test

import (
	"fmt"
	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"testing"
)

func init() {
	router := gin.Default() // 这需要写到init中，启动gin框架
	router.POST("/login", LoginHandler)
	utils.SetRouter(router) //把启动的engine 对象传入到test框架中
}

// 解析返回的错误类型
type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

// 真正的测试单元
func TestLoginHandler(t *testing.T) {
	// 定义发送POST请求传的内容
	user := map[string]interface{}{
		"username": "admin123",
		"password": "123456",
		"age":      13,
	}
	// 把返回response解析到resp中
	resp := OrdinaryResponse{}
	// 调用函数发起http请求
	err := utils.TestHandlerUnMarshalResp("POST", "/login", "json", user, &resp)
	if err != nil {
		t.Errorf("TestLoginHandler: %v\n", err)
		return
	}
	// 得到返回数据结构体， 至此，完美完成一次post请求测试，
	// 如果需要benchmark 输出性能报告也是可以的
	fmt.Println("result:", resp)
}

//[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
//
//[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
//- using env:	export GIN_MODE=release
//- using code:	gin.SetMode(gin.ReleaseMode)
//
//[GIN-debug] POST   /login                    --> gin-leaning-project/Test.LoginHandler (3 handlers)
//=== RUN   TestLoginHandler
//[GIN] 2022/10/09 - 20:34:17 | 200 |      1.0476ms |                 | POST     "/login"
//result: {0 login success}
//--- PASS: TestLoginHandler (0.01s)
//PASS
