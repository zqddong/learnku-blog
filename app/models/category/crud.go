package category

import (
	"github.com/zqddong/learnku-blog/pkg/logger"
	"github.com/zqddong/learnku-blog/pkg/model"
)

// Create 创建分类，通过 category.ID 来判断是否创建成功
func (c *Category) Create() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
