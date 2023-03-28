package articles

import (
	"github.com/zqddong/learnku-blog/pkg/model"
	"github.com/zqddong/learnku-blog/pkg/types"
)

func Get(idStr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idStr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}
