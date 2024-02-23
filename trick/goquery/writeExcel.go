package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsxNew := excelize.NewFile()
	newSheet := "Sheet1"
	xlsxNew.NewSheet(newSheet)

	result := [5][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	for i, row := range result {
		fmt.Printf("(=)i:%d ", i)
		for j, colCell := range row {
			fmt.Printf("j:%d ", j)
			zuobiao := ChangIndexToAxis(i, j)
			ModifyExcelCellByAxis(xlsxNew, newSheet, zuobiao, colCell)
		}
	}
	xlsxNew.SaveAs("./你要的表格.xlsx")
}

// 数组下标转换成excel坐标
func ChangIndexToAxis(intIndexX int, intIndexY int) string {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	intIndexY = intIndexY + 1
	resultY := ""
	for true {
		if intIndexY <= 26 {
			resultY = resultY + arr[intIndexY-1]
			break
		}
		mo := intIndexY % 26
		resultY = arr[mo-1] + resultY
		shang := intIndexY / 26
		if shang <= 26 {
			resultY = arr[shang-1] + resultY
			break
		}
		intIndexY = shang
	}
	return resultY + strconv.Itoa(intIndexX+1)
}

// 修改excel表格里的值
func ModifyExcelCellByAxis(xlsx *excelize.File, sheet string, axis string, value interface{}) int {
	xlsx.SetCellValue(sheet, axis, value)
	return 0
}
