package main

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/lunny/log"
	"github.com/tealeg/xlsx"
)

type TestExcel struct {
	Name   string `TExcel:"姓名"`
	Age    int    `TExcel:"年龄"`
	Gender string `TExcel:"性别"`
}

func main() {

	list := []TestExcel{{"aaa", 28, "male"}, {"bbb", 26, "female"}}

	var interfaceSlice = make([]interface{}, len(list))
	for k, v := range list {
		interfaceSlice[k] = v
	}

	err := GenerateExcelByStruct(interfaceSlice, "test.xlsx")
	if err != nil {
		log.Errorf("GenerateExcelByStruct err:%s\n", err)
	}
}

func GenerateExcelByStruct(r []interface{}, filename string) error {
	var xls [][]interface{}
	xls = make([][]interface{}, 0)

	if len(r) == 0 {
		return errors.New("list is empty")
	}

	titles, err := getExcelTitle(r[0])
	if err != nil {
		return err
	}
	var interfaceSlice = make([]interface{}, len(titles))
	for i, d := range titles {
		interfaceSlice[i] = d
	}
	xls = append(xls, interfaceSlice)

	for i := 0; i < len(r); i++ {
		vs, err := getFieldByStruct(r[i])
		if err != nil {
			return err
		}
		xls = append(xls, vs)
	}

	excel, err := NewExcel()
	if err != nil {
		log.Errorf("NewExcel err:%s\n", err)
		return err
	}

	for _, xl := range xls {
		excel.Write(xl)
	}

	excelFilePath := "/home/lizhenfeng/github/go/" + filename

	err = excel.Save(excelFilePath)
	if err != nil {
		log.Errorf("save file err:%s\n", err)
		return err
	}

	return nil
}

// 一次写一行
func (e *Excel) Write(values []interface{}) {
	row := e.Sheet.AddRow()
	for _, cellValue := range values {
		cell := row.AddCell()
		cell.SetValue(cellValue)
	}
}

// 保存
func (e *Excel) Save(path string) error {
	err := e.File.Save(path)
	if err != nil {
		return err
	}
	return nil
}

func getExcelTitle(obj interface{}) (ret []string, err error) {
	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		err = errors.New("no struct")
		return
	}

	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	ret = make([]string, 0)

	TExcel := ""
	for i := 0; i < v.NumField(); i++ {
		TExcel = t.Field(i).Tag.Get("TExcel")
		if len(TExcel) > 0 {
			ret = append(ret, TExcel)
		}
	}
	return
}

func getFieldByStruct(structName interface{}) ([]interface{}, error) {
	ret := make([]interface{}, 0)
	TValue := ""
	v := reflect.ValueOf(structName)
	if !v.IsValid() {
		log.Error("no struct")
		return nil, errors.New("no struct")
	}

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		kind := value.Kind()
		switch kind {
		case reflect.String:
			TValue = value.String()
		case reflect.Int:
			TValue = strconv.FormatInt(value.Int(), 10)
		}
		ret = append(ret, TValue)
	}
	return ret, nil
}

type Excel struct {
	File  *xlsx.File
	Sheet *xlsx.Sheet
}

func NewExcel() (*Excel, error) {
	excel := new(Excel)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}
	excel.File = file
	excel.Sheet = sheet

	return excel, nil
}
