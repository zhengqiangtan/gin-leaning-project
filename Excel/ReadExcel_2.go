package main

import (
	"fmt"
	"github.com/tealeg/xlsx/v2" //注意 v2版本的简单一些
)

func main() {
	ReadExcel()
}

type Person struct {
	Name       string // 姓名
	Education  string // 学历
	University string // 高校
	Industry   string // 行业
	Workyear   string // 工作年限
	Position   string // 职位
	Salary     string // 薪资
	Language   string // 编程语言
}

func ReadExcel() {
	var per []Person

	// 打开 xlsx文件
	xlFile, err := xlsx.OpenFile("ccmous.xlsx")
	if err != nil {
		fmt.Println("打开文件失败", err.Error())
		return
	}

	// 遍历 sheet 页
	for _, sheet := range xlFile.Sheets {
		// 行
		for _, row := range sheet.Rows {
			// 列
			var temp Person
			// 将excel每一列文件读取放在字符串切片中
			var str []string
			for _, cell := range row.Cells {
				str = append(str, cell.String())
			}
			temp.Name = str[0]
			temp.Education = str[1]
			temp.University = str[2]
			temp.Industry = str[3]
			temp.Workyear = str[4]
			temp.Position = str[5]
			temp.Salary = str[6]
			temp.Language = str[7]
			per = append(per, temp)
		}

	}
	fmt.Println(per)
	// db
}
