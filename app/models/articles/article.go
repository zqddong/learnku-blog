package articles

import (
	"github.com/zqddong/learnku-blog/app/models"
	"github.com/zqddong/learnku-blog/app/models/user"
	"github.com/zqddong/learnku-blog/pkg/route"
	"strconv"
)

type Article struct {
	models.BaseModel

	Title  string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body   string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}

// CreatedAtDate 创建日期
func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
