package category

import (
	"github.com/zqddong/learnku-blog/pkg/logger"
	"github.com/zqddong/learnku-blog/pkg/model"
	"github.com/zqddong/learnku-blog/pkg/types"
)

// Create 创建分类，通过 category.ID 来判断是否创建成功
func (c *Category) Create() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

// Get 通过 ID 获取分类
func Get(idStr string) (Category, error) {
	var category Category
	id := types.StringToUint64(idStr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}
