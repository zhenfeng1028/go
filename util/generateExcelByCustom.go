package main

import (
	"errors"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/lunny/log"
	"github.com/tealeg/xlsx"
)

type Person struct {
	Name   string `TExcel:"姓名"`
	Age    int    `TExcel:"年龄"`
	Gender string `TExcel:"性别"`
}

func main() {
	list := []Person{{"aaa", 28, "male"}, {"bbb", 26, "female"}}

	interfaceSlice := make([]interface{}, len(list))
	for k, v := range list {
		interfaceSlice[k] = v
	}

	requiredFields := []interface{}{"姓名", "年龄"}
	err := GenerateExcelByCustom(interfaceSlice, "person.xlsx", requiredFields)
	if err != nil {
		log.Errorf("GenerateExcelByCustom err: %s\n", err)
	}
}

func GenerateExcelByCustom(r []interface{}, filename string, requiredFields []interface{}) error {
	var xls [][]interface{}
	xls = make([][]interface{}, 0)

	if len(r) == 0 {
		return errors.New("list is empty")
	}

	titles, err := getExcelTitle(r[0], requiredFields)
	if err != nil {
		return err
	}
	interfaceSlice := make([]interface{}, len(titles))
	copy(interfaceSlice, titles)
	xls = append(xls, interfaceSlice)

	for i := 0; i < len(r); i++ {
		vs, err := getFieldByStruct(r[i], requiredFields)
		if err != nil {
			return err
		}
		xls = append(xls, vs)
	}

	excel, err := NewExcel()
	if err != nil {
		log.Errorf("NewExcel err: %s\n", err)
		return err
	}

	for _, xl := range xls {
		excel.Write(xl)
	}

	excelFilePath := "./../assets/" + filename
	excelFilePath, err = filepath.Abs(excelFilePath)
	if err != nil {
		log.Errorf("filepath.Abs err: %s\n", err)
		return err
	}

	err = excel.Save(excelFilePath)
	if err != nil {
		log.Errorf("save file err: %s\n", err)
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

type ExcelField struct {
	TName  string
	TValue interface{}
}

func getExcelTitle(obj interface{}, requiredFields []interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(obj)
	if !v.IsValid() {
		return nil, errors.New("no struct")
	}

	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	prepare := make([]interface{}, 0)
	for i := 0; i < v.NumField(); i++ {
		TName := t.Field(i).Tag.Get("TExcel")
		if len(TName) > 0 {
			prepare = append(prepare, ExcelField{TName, TName})
		}
	}

	ret := FilterFields(prepare, requiredFields)
	return ret, nil
}

func getFieldByStruct(structName interface{}, requiredFields []interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(structName)
	if !v.IsValid() {
		return nil, errors.New("no struct")
	}

	prepare := make([]interface{}, 0)
	for i := 0; i < v.NumField(); i++ {
		TValue := ""
		TName := v.Type().Field(i).Tag.Get("TExcel")
		value := v.Field(i)
		kind := value.Kind()
		switch kind {
		case reflect.String:
			TValue = value.String()
		case reflect.Int:
			TValue = strconv.FormatInt(value.Int(), 10)
		}
		prepare = append(prepare, ExcelField{TName, TValue})
	}

	ret := FilterFields(prepare, requiredFields)
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

func FilterFields(fields []interface{}, requiredFields []interface{}) []interface{} {
	tmp := make([]interface{}, 0)
	if len(fields) > 0 && len(requiredFields) > 0 {
		locFields := make([]ExcelField, 0)
		for _, v := range fields {
			locFields = append(locFields, v.(ExcelField))
		}
		locRequiredFields := make([]string, 0)
		for _, v := range requiredFields {
			locRequiredFields = append(locRequiredFields, v.(string))
		}

		for _, r := range locFields {
			for _, n := range locRequiredFields {
				if r.TName == n {
					tmp = append(tmp, r.TValue)
					break
				}
			}
		}
	}

	return tmp
}
