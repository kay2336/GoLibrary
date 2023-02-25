package Main

import (
	"fmt"
	"goGormLibrary/curd"
	"goGormLibrary/mysql"
)

func Menu() {
	//初始化连接 MySQL
	mysql.ConnMysql()

	//输出主菜单
	fmt.Println("1. CreatOne")
	fmt.Println("2. DisplayAll")
	fmt.Println("3. SearchOne")
	fmt.Println("4. DeleteOne")
	fmt.Println("5. ModifyOne")
	var choice int
	for choice != -1 {
		//更新数据库数据至本地
		mysql.UpdateConn()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			curd.CreatOne()
		case 2:
			curd.DisplayAll()
		case 3:
			curd.SearchOne()
		case 4:
			curd.DeleteOne()
		case 5:
			curd.ModifyOne()
		default:
			fmt.Println("无效指令")
			fmt.Println("已退出图书管理系统")
			choice = -1
		}
	}
}
