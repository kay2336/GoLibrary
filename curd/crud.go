package curd

import (
	"fmt"
	"goGormLibrary/model"
	"goGormLibrary/mysql"
)

// CreatOne
// 新增图书至数据库
func CreatOne() {
	//输出新建图书的提示信息
	fmt.Println("正在新增图书")

	//输入新增图书的基本信息
	var book model.Book
	fmt.Scan(&book.Title, &book.Author)

	//有相同的Id或Title则输出提示信息，没有则新增至数据库
	if mysql.LibraryCreatOne(book) {
		fmt.Println("新建图书成功")
	} else {
		fmt.Println("新增图书失败，原因是数据库中已有相同编号或者书名相同的图书")
	}
}

// DisplayAll
// 查询数据库中所有的图书
func DisplayAll() {
	//输出查询所有图书的提示信息
	fmt.Println("正在查询所有图书")

	//查询功能
	mysql.LibraryDisplayAll()

	//对查询所有图书操作输出对应提示信息
	if len(model.BookSlice) > 0 {
		fmt.Printf("已查询所有图书, 一共%v本图书\n", len(model.BookSlice))
	} else {
		fmt.Println("未能在数据库中查询到此图书")
	}
}

// SearchOne
// 查询数据库中某本图书
func SearchOne() {
	//输出查询某本图书的提示信息
	fmt.Println("请输入需要查询图书的书名")

	//查询功能
	fmt.Scan(&model.Title)
	tBookSlice := mysql.LibrarySearchOne()

	//对查询某本图书操作输出对应提示信息
	if tBookSlice == nil {
		fmt.Println("未能在数据库中查询到此图书")
	} else {
		fmt.Println(tBookSlice)
		fmt.Println("查询图书成功")
	}
}

// DeleteOne
// 删除数据库中某本图书
func DeleteOne() {
	//输出删除界面的提示信息
	fmt.Println("请输入需要删除图书的书名")
	fmt.Scan(&model.Title)

	//删除功能及对删除操作输出对应提示信息
	if mysql.LibraryDeleteOne(model.Title) {
		fmt.Println("删除图书成功")
	} else {
		fmt.Println("未能在数据库中找到此图书")
	}
}

// ModifyOne
// 修改数据库中图书的信息
func ModifyOne() {
	//输出进入修改界面的提示信息
	fmt.Println("请输入需要修改图书的书名")
	fmt.Scan(&model.Title)

	//修改功能及对删除操作输出对应提示信息
	if mysql.LibraryModifyOne(model.Title) {
		fmt.Println("修改图书成功")
	} else {
		fmt.Println("未能在数据库中找到此图书")
	}
}
