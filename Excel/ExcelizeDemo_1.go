package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// github.com/xuri/excelize
// docs: https://xuri.me/excelize/zh-hans/
func main() {
	f, err := excelize.OpenFile("C:\\Users\\HNK7WC3\\Downloads\\Feedback_221011180332.xlsx")
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
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	display, tooltip := "https://github.com/xuri/excelize", "Excelize on GitHub"
	if err := f.SetCellHyperLink("Sheet1", "B2",
		"https://github.com/xuri/excelize", "External", excelize.HyperlinkOpts{
			Display: &display,
			Tooltip: &tooltip,
		}); err != nil {
		fmt.Println(err)
	}

	for i, row := range rows {
		var str []string
		for _, colCell := range row {
			str = append(str, colCell)
		}
		if i == 2 {
			break
		}
		fmt.Println(str)
	}
	f.Save()
}
