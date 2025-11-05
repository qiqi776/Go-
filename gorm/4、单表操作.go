package main

import (
	"fmt"
	"gorm-project/global"
	"gorm-project/models"
	"time"
)

// 插入
func insert() {
	err := global.DB.Create(&models.UserModel{
		Name: "柒风2",
		Age:  18,
	}).Error
	if err != nil {
		fmt.Println(err)
	}

	//回填式创建
	user := models.UserModel{
		Name: "张三2",
		Age:  18,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.ID, user.Name, user.Age, user.CreatedAt)

	var userList = []models.UserModel{
		{Name: "王五"},
		{Name: "李四"},
	}
	err = global.DB.Create(&userList).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userList)

	err = global.DB.Create(&models.UserModel{
		Name: "柒风3",
		Age:  18,
	}).Error
	if err != nil {
		fmt.Println(err)
	}
}

// 查询
func query() {
	var userList []models.UserModel
	global.DB.Find(&userList, "age > 18")
	fmt.Println(userList)

	// var user models.UserModel
	// global.DB.Take(&user)
	// 根据主键查询
	// global.DB.Take(&user, 4)

	// 根据主键排序查第一个
	// Debug能够生成对应的sql
	//global.DB.Debug().First(&user, 4)
	//fmt.Println(user)
	//global.DB.Debug().Last(&user, 5)
	//fmt.Println(user)

	// 使用limit，即便查找不到也不会报错
	//var user models.UserModel
	//err := global.DB.Limit(1).Find(&user, 111).Error
	//if err != gorm.ErrRecordNotFound {
	//	fmt.Println("记录不存在")
	//	return
	//}
	//fmt.Println(user)
}

// 更新
func updateSave() {
	var user models.UserModel
	user.ID = 4
	// 会更新零值
	user.Name = "柒风4"
	user.CreatedAt = time.Now()
	global.DB.Save(&user)
	fmt.Println(user)
}

func update() {
	var user = models.UserModel{ID: 3}
	global.DB.Model(&user).Update("age", 34)
	// global.DB.Model(&user).UpdateColumn("age", 34)
	fmt.Println(user)
}

func updates() {
	var user = models.UserModel{ID: 3}
	//global.DB.Model(&user).Updates(models.UserModel{
	//	Name: "zhangsan",
	//	// 不会更新零值
	//	Age: 22,
	//}

	global.DB.Model(&user).Updates(map[string]any{
		// 可以更新零值
		"age": 0,
	})
	fmt.Println(user)
}

// 删除
func Delete() {
	var user = models.UserModel{ID: 3}
	//global.DB.Delete(&user)
	//global.DB.Delete(&models.UserModel{}, 8)
	global.DB.Delete(&models.UserModel{}, []int{4, 5})
	fmt.Println(user)
}

// 软删除

func main() {
	global.Connect()
	// query()
	// updateSave()
	// update()
	// updates()
	Delete()
}
