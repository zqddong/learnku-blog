package policies

import (
	"github.com/zqddong/learnku-blog/app/models/articles"
	"github.com/zqddong/learnku-blog/pkg/auth"
)

// CanModifyArticle 是否允许修改话题
func CanModifyArticle(_article articles.Article) bool {
	return auth.User().ID == _article.UserID
}
