package user

import (
	"github.com/zqddong/learnku-blog/app/models"
	"github.com/zqddong/learnku-blog/pkg/password"
)

type User struct {
	models.BaseModel

	//GORM 默认会将键小写化作为字段名称，column 项可去除，另外默认是允许 NULL 的，故 default:NULL 项也可去除
	//Name     string `gorm:"column:name;type:varchar(255);not null;unique"`
	//Email    string `gorm:"column:email;type:varchar(255);default:NULL;unique;"`
	//Password string `gorm:"column:password;type:varchar(255)"`
	//// gorm:"-" —— 设置 GORM 在读写时略过此字段
	//PasswordConfirm string `gorm:"-" valid:"password_confirm"`

	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// ComparePassword 对比密码是否匹配
func (user *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, user.Password)
}

// Link 方法用来生成用户链接
func (user *User) Link() string {
	return ""
}
