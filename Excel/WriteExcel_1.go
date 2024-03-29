package main

import (
	"fmt"
	"gin-leaning-project/globals"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx/v3"
	"github.com/wxnacy/wgo/arrays"
	"log"
	"net/http"
	"time"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type FeedbackReplyModelDto struct {
	Id           string `json:"id"`
	App          string `json:"app"`
	Modular      string `json:"modular"`
	QuestionDesc string `json:"question_desc"`
	ContentEn    string `json:"content_en"`
	ContentCn    string `json:"content_cn"`
}

func UNIONRoutes(route *gin.Engine) {
	UnionRouter := route.Group("/union")
	exportRouter := UnionRouter.Group("/export")
	exportRouter.GET("/reply", HandleExportReply2Excel)
}

// 导出excel测试
func main() {
	gin.SetMode("debug")
	r := gin.Default()
	UNIONRoutes(r)
	r.Run(":8080")
	// curl --location --request GET 'http://127.0.0.1:8080/union/export/reply' --output reply1.xlsx
}

func HandleExportReply2Excel(c *gin.Context) {
	err, tbList := GetReplyTable()
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{
			Code:    globals.ErrDbQueryException.Code,
			Message: globals.ErrDbQueryException.Message,
			Data:    err,
		})
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf(`attachment; filename="reply_%s.xlsx"`, time.Now().Format("060102150405"))
	c.Header("Content-Disposition", disposition)

	ExportFeedbackReplyExcel(tbList, c)
}
func GetReplyTable() (error, []FeedbackReplyModelDto) {
	var list = make([]FeedbackReplyModelDto, 0, 1)
	// db select ...
	list = append(list, FeedbackReplyModelDto{
		Id:           "10",
		App:          "android",
		Modular:      "pub",
		QuestionDesc: "test",
		ContentEn:    "this is a test",
		ContentCn:    "这是个测试",
	})
	list = append(list, FeedbackReplyModelDto{
		Id:           "11",
		App:          "ios",
		Modular:      "pub",
		QuestionDesc: "test",
		ContentEn:    "taa",
		ContentCn:    "fdsf",
	})

	return nil, list
}

func ExportFeedbackReplyExcel(fbs []FeedbackReplyModelDto, c *gin.Context) {
	file := xlsx.NewFile()

	sheet, _ := file.AddSheet("Sheet1")
	sheet.SetColWidth(3, 10, 12.5)

	title := []string{"app", "modular", "question_desc", "content_en", "content_cn"}
	titleRow2 := sheet.AddRow()

	myStyle := xlsx.NewStyle()
	myStyle.Alignment.Horizontal = "center"
	myStyle.Fill.FgColor = "FFFFFF00"
	myStyle.Fill.PatternType = "solid"
	myStyle.Font.Name = "Georgia"
	myStyle.Font.Size = 13

	myStyle.Font.Bold = true
	myStyle.ApplyAlignment = true
	myStyle.ApplyFill = true
	myStyle.ApplyFont = true

	for _, title := range title {
		cell := titleRow2.AddCell()
		cell.Value = title
		cell.SetStyle(myStyle)
	}
	//测试动态添加表格
	colsEn := []string{"created_at", "feedback_id"}
	for _, fb := range fbs {
		row := sheet.AddRow()
		if arrays.ContainsString(colsEn, "created_at") > 0 {
			cell := row.AddCell()
			cell.Value = fb.App
		}
		if arrays.ContainsString(colsEn, "feedback_id") > 0 {
			cel := row.AddCell()
			cel.Value = fb.Modular
			cel.SetHyperlink("https://feedback-admin.4wps.net/", fb.App, "")
			cel.GetStyle().Font.Color = "0B33F4"

		}
		cell3 := row.AddCell()
		cell3.Value = fb.QuestionDesc
		cell4 := row.AddCell()
		cell4.Value = fb.ContentEn
		cell5 := row.AddCell()
		cell5.Value = fb.ContentCn

	}
	err := file.Write(c.Writer)
	if err != nil {
		log.Println("save file fail:", err)
	}
}
