package model

import "shareboard/util"

//执行数据迁移

func migration() {
	// 自动迁移模式
	if err := DB.AutoMigrate(&User{}); err != nil {
		util.Log().Error("AutoMigrate User Failed", err)
	}
	if err := DB.AutoMigrate(&Page{}); err != nil {
		util.Log().Error("AutoMigrate Page Failed", err)
	}
}
