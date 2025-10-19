package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Student{})

	// 1.编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	result := db.Create(&Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	})

	if result.Error != nil {
		fmt.Printf("插入失败 %v\n", result.Error)
	}

	// 2.编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	for _, student := range students {
		fmt.Printf("查询的结果%v \n", student)
	}

	// 3.编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	updateResult := db.Model(&Student{}).
		Where("name =?", "张三").
		Update("grade", "四年级")

	if updateResult.Error != nil {
		fmt.Printf("更新失败 %v\n", updateResult.Error)
	}

	// 3.编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	deleteResult := db.Where("age < ?", 15).Delete(&Student{})
	if deleteResult.Error != nil {
		fmt.Printf("删除失败,%v", deleteResult.Error)
	}

	var all []Student
	db.Find(&all)
	for _, student := range all {
		fmt.Printf("id : %d,name: %s,age: %d,grade : %s \n", student.ID, student.Name, student.Age, student.Grade)
	}

}

type Student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}
