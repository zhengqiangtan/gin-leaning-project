package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HMACSHA256加密-hex格式且转大写
func HmacSha256ForHexUpper(message string, secretKey string) string {
	sign := HmacSha256ForHex(message, secretKey)
	return strings.ToUpper(sign)
}

// HMACSHA256加密-hex格式
func HmacSha256ForHex(message string, secretKey string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	m := make(map[string]interface{}, 4)
	m["app"] = "wps"
	m["feedback_type"] = "auto"
	m["question_type"] = "uninstall"
	m["question_type"] = "uninstall"
	m["content"] = "test content"

	var arr = make([]interface{}, 0, 10)
	n := make(map[string]interface{}, 5)
	n["attach_type"] = "video"
	n["filename"] = "example-filename.mp4"
	n["cs_provider"] = "s3://example-bucket"
	n["cs_url"] = "http://example-bucket.s3.exmaple.com/example-prefix/example-filename.mp4"
	n["size"] = 8516446546
	arr = append(arr, n)
	m["attachs"] = arr

	jsonBody, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//测试反序列化
	req1 := ReportRequest{}
	json.Unmarshal([]byte(jsonBody), &req1)
	fmt.Println("result = ", string(jsonBody))

	ts := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println("ts===" + ts)
	var str = string(jsonBody) + "&timestamp=" + ts

	fmt.Println(str)
	auth := HmacSha256ForHexUpper(str, "SK")
	fmt.Println(auth)

	//Golang发送post请求
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", "http://feedback-us-test.4wps.net/union/frontapi/api/sign/report", bytes.NewReader(jsonBody))
	//req, err := http.NewRequest("POST", "http://localhost:8081/api/sign/report", bytes.NewReader(jsonBody))
	if err != nil {
		fmt.Println(err.Error())
		//return fmt.Errorf("Got error %s", err.Error())
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("AccessKey", "AK")
	req.Header.Add("Authorization", auth)
	req.Header.Add("Timestamp", ts)
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//defer response.Body.Close()

}

type ReportAttach struct {
	AttachType string `json:"attach_type" binding:"required,max=32"`
	Filename   string `json:"filename" binding:"required"`
	CSProvider string `json:"cs_provider" binding:"required"`
	CSURL      string `json:"cs_url" binding:"required"`
	Size       int64  `json:"size" binding:"required"`
}

type ReportRequest struct {
	App          string `json:"app" binding:"required,min=1,max=32"`            // app代码，wps（wps系列）， docs（在线文档系列），pdf（pdf系列）
	FeedbackType string `json:"feedback_type"  binding:"required,min=1,max=32"` // 反馈分类，auto, feedback, auto为自动上报事件，如卸载上报等，feedback为用户反馈
	QuestionType string `json:"question_type" binding:"required,min=1,max=128"` // 问题类型，用户反馈为具体的类型：好评， 功能建议，登录问题等等，自动上报为对应的上报类别：卸载，崩溃等
	// 主动反馈用字段， 不用时留空
	Content     string         `json:"content"  binding:"max=65536"`  // 反馈内容
	ContactType string         `json:"contact_type" binding:"max=32"` // 联系方式类型 `email`,`mobile`
	Contact     string         `json:"contact" binding:"max=256"`     // 联系方式内容
	Attachs     []ReportAttach `json:"attachs"`                       // 上传文件信息
	// 端信息
	ClientType          string `json:"client_type" binding:"max=32"`           // 客户端类型， pc, android, web, mac, ios, linux
	OSVersion           string `json:"os_version" binding:"max=32"`            // 系统版本， android为系统版本 10， 11，pc为 10.0 , 6.3 , 6.2; web 留空
	Version             string `json:"version" binding:"max=32"`               // 客户端版本， 各个客户端定义，与参数系统的规则保持一致
	Channel             string `json:"channel" binding:"max=32"`               // 渠道， 渠道号， web留空
	ClientLang          string `json:"client_lang" binding:"max=32"`           // 客户端选定语言，指用户选定的语言，没选定时与系统一致，无多语言的应用应与端开发语言一致，如只有英文版的应用ClientLang应为en-US
	SysLang             string `json:"sys_lang" binding:"max=32"`              // 系统语言, web为浏览器默认语言第一项
	SysDesc             string `json:"sys_desc"  binding:"max=65536"`          // 系统描述， pc为设备信息，android为设备名， web为浏览器信息
	NetStat             string `json:"net_stat" binding:"max=128"`             // 网络状态
	VipType             string `json:"vip_type" binding:"max=32"`              // 用户VIP类型，会员，注册用户，游客
	UserID              string `json:"user_id" binding:"max=32"`               // 用户ID，为登录留空
	LicenseStatus       string `json:"license_status" binding:"max=32"`        // license状态，激活，为激活
	Modular             string `json:"modular" binding:"max=128"`              // 组件，'公共，演示，文字，表格'
	EntranceName        string `json:"entrance_name" binding:"max=128"`        // 入口名
	Brand               string `json:"brand" binding:"max=128"`                // 设备品牌，android用， ios， mac固定为apple，pc和web留空备用
	Browser             string `json:"browser" binding:"max=128"`              // 浏览器品牌，chrome，edge，firefox等，web用，其他留空
	BrowserVersion      string `json:"browser_version" binding:"max=32"`       // 浏览器版本
	BrowserCore         string `json:"browser_core" binding:"max=128"`         // 浏览器内核，webkit，geetoo等
	FirstInstallWpsTime int64  `json:"first_install_wps_time" binding:"min=0"` // 首次安装wps的unix时间(单位秒), pc用，其余留0
	Events              string `json:"events" binding:"max=65536"`             // 埋点信息,暂时留空
	CustomParam         string `json:"custom_param"`                           // 其他信息
	DeviceId            string `json:"device_id"`                              //设备id
	IP                  string `json:"-"`
	MockIP              string `json:"mock_ip"`
	UserAgent           string `json:"-"`
	AdPosition          string `json:"ad_position"`
	AdReason            string `json:"ad_reason"`
	AdPlatform          string `json:"ad_platform"`
	AdMaterialPicture   string `json:"ad_material_picture"`
	AdMaterialContent   string `json:"ad_material_content"`
	WpsAdId             string `json:"wps_ad_id"`   // 广告id
	DataSource          string `json:"data_source"` // 来源
}
