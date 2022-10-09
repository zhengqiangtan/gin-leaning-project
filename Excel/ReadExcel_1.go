package main

import "fmt"
import "github.com/tealeg/xlsx/v3"

// https://github.com/tealeg/xlsx/blob/master/tutorial/tutorial.adoc#getting-started
// Doing
func main() {
	wb, err := xlsx.OpenFile("C:\\Users\\HNK7WC3\\Downloads\\samplefile.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name) //打印工作表名称
	}
	fmt.Println("----")

	fmt.Println("Workbook contains", len(wb.Sheets), "sheets.")

	sheetName := "Sample"
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	fmt.Println("Max row in sheet:", sh.MaxRow)

}
