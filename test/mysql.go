package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Student struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	No       string `json:"no" xorm:"not null comment('学号') VARCHAR(255)"`
	Name     string `json:"name" xorm:"not null comment('姓名') VARCHAR(255)"`
	Sex      string `json:"sex" xorm:"not null comment('性别') VARCHAR(255)"`
	Birthday string `json:"birthday" xorm:"comment('生日') VARCHAR(255)"`
	Class    string `json:"class" xorm:"comment('班级') VARCHAR(255)"`
}

type Teacher struct {
	Id         int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	No         string `json:"no" xorm:"not null comment('工号') VARCHAR(255)"`
	Name       string `json:"name" xorm:"not null comment('姓名') VARCHAR(255)"`
	Sex        string `json:"sex" xorm:"not null comment('性别') VARCHAR(255)"`
	Birthday   string `json:"birthday" xorm:"comment('生日') VARCHAR(255)"`
	Title      string `json:"title" xorm:"not null comment('职称') VARCHAR(255)"`
	Department string `json:"department" xorm:"not null comment('部门') VARCHAR(255)"`
}

type Course struct {
	Id   int    `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	No   string `json:"no" xorm:"not null comment('课程编号') VARCHAR(255)"`
	Name string `json:"name" xorm:"not null comment('课程名') VARCHAR(255)"`
	Tno  string `json:"tno" xorm:"not null comment('教师编号') VARCHAR(255)"`
}

type Score struct {
	Id    int     `json:"id" xorm:"not null pk autoincr comment('自增id') INT(11)"`
	Sno   string  `json:"sno" xorm:"not null comment('学号') VARCHAR(255)"`
	Cno   string  `json:"cno" xorm:"not null comment('课程编号') VARCHAR(255)"`
	Grade float64 `json:"grade" xorm:"comment('成绩') FLOAT"`
}

func main() {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:smai123@(100.100.142.15:4000)/lzf?charset=utf8") // lzf是数据库实例名称
	if err != nil {
		log.Fatal("init engine err: ", err)
	}
	if err := engine.Sync2(new(Student), new(Teacher), new(Course), new(Score)); err != nil {
		log.Fatal("sync database err: ", err)
	}

	session := engine.NewSession()
	defer session.Close()

	student := []Student{
		{No: "101", Name: "曾华", Sex: "男", Birthday: "1977-09-01", Class: "95033"},
		{No: "102", Name: "匡明", Sex: "男", Birthday: "1975-10-02", Class: "95031"},
		{No: "103", Name: "王丽", Sex: "女", Birthday: "1976-01-23", Class: "95033"},
		{No: "104", Name: "李军", Sex: "男", Birthday: "1976-02-20", Class: "95033"},
		{No: "105", Name: "王芳", Sex: "女", Birthday: "1975-02-10", Class: "95031"},
		{No: "106", Name: "陆军", Sex: "男", Birthday: "1974-06-03", Class: "95031"},
		{No: "107", Name: "王尼玛", Sex: "男", Birthday: "1976-02-20", Class: "95033"},
		{No: "108", Name: "张全蛋", Sex: "男", Birthday: "1975-02-10", Class: "95031"},
		{No: "109", Name: "赵铁柱", Sex: "男", Birthday: "1974-06-03", Class: "95031"},
	}

	teacher := []Teacher{
		{No: "804", Name: "李诚", Sex: "男", Birthday: "1958-12-02", Title: "副教授", Department: "计算机系"},
		{No: "856", Name: "张旭", Sex: "男", Birthday: "1969-03-12", Title: "讲师", Department: "电子工程系"},
		{No: "825", Name: "王萍", Sex: "女", Birthday: "1972-05-05", Title: "助教", Department: "计算机系"},
		{No: "831", Name: "刘冰", Sex: "女", Birthday: "1977-08-14", Title: "副教授", Department: "电子工程系"},
	}

	course := []Course{
		{No: "3-105", Name: "计算机导论", Tno: "825"},
		{No: "3-245", Name: "操作系统", Tno: "804"},
		{No: "6-166", Name: "数字电路", Tno: "856"},
		{No: "9-888", Name: "高等数学", Tno: "831"},
	}

	score := []Score{
		{Sno: "103", Cno: "3-105", Grade: 92},
		{Sno: "103", Cno: "3-245", Grade: 86},
		{Sno: "103", Cno: "6-166", Grade: 85},
		{Sno: "105", Cno: "3-105", Grade: 88},
		{Sno: "105", Cno: "3-245", Grade: 75},
		{Sno: "105", Cno: "6-166", Grade: 79},
		{Sno: "109", Cno: "3-105", Grade: 76},
		{Sno: "109", Cno: "3-245", Grade: 68},
		{Sno: "109", Cno: "6-166", Grade: 81},
	}

	_, err = session.Insert(student, teacher, course, score)
	if err != nil {
		log.Println(err)
	}
}
