package mysql

import (
	"errors"
	"fmt"
	"goGormLibrary/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnMysql
// 连接数据库
func ConnMysql() {
	// gorm 连接到 mysql， 获得 db 实例
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_gorm_library?" +
		"charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	model.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 自动创建表
	model.DB.AutoMigrate(&model.Book{})
}

// UpdateConn
// 将数据库的数据导入结构体切片内
func UpdateConn() {
	model.DB.Find(&model.BookSlice)
}

// LibraryCreatOne
// 将图书的基本信息存入数据库
// 注意：此处的写法会出现record not found的错误信息，但并不影响新增功能的实现
// 因为此错误信息指的是未能在数据库查询到包含给定书名和作者的图书
// record not found：说明查询结果为空，没有找到符合条件的记录
func LibraryCreatOne(book model.Book) bool {
	//遍历数据库，查找是否存在Id或者Title相同的图书
	var tBookSlice []model.Book
	//得到错误信息（未能在数据库中找到与新增图书相同信息的图书）
	err := model.DB.Where("Title = ? AND Author = ?",
		book.Title, book.Author).First(&tBookSlice).Error
	//判断err是否为gorm.ErrRecordNotFound类型(一种特殊的错误类型，用于表示查询记录不存在的情况，其底层实现是一个字符串常量)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//向数据库中新增图书
		model.DB.Create(&model.Book{
			Title:  book.Title,
			Author: book.Author,
		})
		return true
	}
	return false

	//等效遍历代码（无错误信息）
	//mark := true
	//for _, v := range model.BookSlice {
	//	if v.Title == book.Title || v.Author == book.Author {
	//		mark = false
	//	}
	//}
	//return true
	//还有通过len(tBookSlice)==0判断是否查询到相同类型的图书
	//上述通过处理错误信息的方式来实现更多的是带有做笔记的含义
}

// LibraryDisplayAll
// 查询数据库中所有的图书
func LibraryDisplayAll() {
	for i := 0; i < len(model.BookSlice); i++ {
		fmt.Println(model.BookSlice[i])
	}
}

// LibrarySearchOne
// 查询数据库中对应书名的图书
func LibrarySearchOne() (tBookSlice []model.Book) {
	// 使用gorm中的Where实现对特定信息的查询
	model.DB.Where(&model.Book{Title: model.Title}).First(&tBookSlice)
	return tBookSlice
}

// LibraryDeleteOne
// 删除数据库中对应书名的图书，数据库中有且删除成功返回true，数据库中没有则返回false
func LibraryDeleteOne(title string) (mark bool) {
	for _, v := range model.BookSlice {
		if v.Title == title {
			//删除数据库中书名为 title 的图书
			model.DB.Where("Title = ?", title).Delete(&model.Book{})
			return true
		}
	}
	return false
}

// LibraryModifyOne
// 修改数据库中对应书名的图书，数据库中有且修改成功返回true，数据库中没有则返回false
func LibraryModifyOne(title string) (mark bool) {
	for _, v := range model.BookSlice {
		if v.Title == title {
			//修改数据库中书名为title的图书
			fmt.Println("请输入修改后的书名")
			fmt.Scan(&model.Title)
			fmt.Println("请输入修改后的作者")
			fmt.Scan(&model.Author)
			//修改书名为title的图书，将其书名修改为model.Title，作者修改为model.Author
			model.DB.Model(&model.Book{}).Where("Title = ?", title).
				Updates(model.Book{Title: model.Title, Author: model.Author})
			return true
		}
	}
	return false
}
