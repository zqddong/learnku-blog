package articles

import (
	"github.com/zqddong/learnku-blog/pkg/logger"
	"github.com/zqddong/learnku-blog/pkg/model"
	"github.com/zqddong/learnku-blog/pkg/types"
)

func Get(idStr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idStr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	//result := model.DB.Save(&article) // 不修改 Save执行 会 UPDATE 后在执行INSERT 没整明白原因
	result := model.DB.Updates(article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}
