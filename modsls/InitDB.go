package modsls

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移
	if err = DB.AutoMigrate(&Articles{}); err != nil {
		return err
	}

	// 检查是否需要添加测试数据
	var count int64
	DB.Model(&Articles{}).Count(&count)
	if count == 0 {
		// 添加测试数据
		testArticles := []Articles{
			{Name: "高级医疗包", Quality: Red, Space: Space{Width: 2, Height: 2}, Weight: 1, Price: 5000, Type: "医疗"},
			{Name: "黄金AK47", Quality: Gold, Space: Space{Width: 4, Height: 2}, Weight: 3, Price: 3500, Type: "武器"},
			{Name: "紫色防弹衣", Quality: Purple, Space: Space{Width: 3, Height: 3}, Weight: 5, Price: 2000, Type: "防具"},
			{Name: "蓝色头盔", Quality: Blue, Space: Space{Width: 2, Height: 2}, Weight: 2, Price: 1000, Type: "防具"},
			{Name: "绿色背包", Quality: Green, Space: Space{Width: 4, Height: 4}, Weight: 2, Price: 500, Type: "装备"},
			{Name: "普通绷带", Quality: White, Space: Space{Width: 1, Height: 1}, Weight: 1, Price: 100, Type: "医疗"},
			{Name: "红色狙击枪", Quality: Red, Space: Space{Width: 5, Height: 2}, Weight: 4, Price: 6000, Type: "武器"},
			{Name: "金色手雷", Quality: Gold, Space: Space{Width: 1, Height: 2}, Weight: 1, Price: 2000, Type: "武器"},
			{Name: "紫色弹药", Quality: Purple, Space: Space{Width: 1, Height: 1}, Weight: 1, Price: 1000, Type: "弹药"},
			{Name: "蓝色饮料", Quality: Blue, Space: Space{Width: 1, Height: 2}, Weight: 1, Price: 300, Type: "消耗品"},
		}

		for _, article := range testArticles {
			if err := DB.Create(&article).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
