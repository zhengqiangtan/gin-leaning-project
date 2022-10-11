package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// github.com/xuri/excelize
// docs: https://xuri.me/excelize/zh-hans/
func main() {
	f, err := excelize.OpenFile("C:\\Users\\HNK7WC3\\Downloads\\反馈系统（标注项）-PC.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("UWP版")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		if i == 2 {
			break
		}
		fmt.Println()
	}
}
