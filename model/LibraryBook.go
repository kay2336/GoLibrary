package model

import "gorm.io/gorm"

// Book 用结构体包含图书的基本信息
type Book struct {
	gorm.Model
	Title  string
	Author string
}

// 定义全局变量
var (
	DB        *gorm.DB
	BookSlice []Book
	Title     string
	Author    string
)
