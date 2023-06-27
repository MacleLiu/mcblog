package modles

import (
	"mcblog/utils/encrypt"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(200);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 查询用户是否存在
func CheckUser(name string) error {
	var user User
	db.Select("id").Where("username=?", name).First(&user)

	if user.ID > 0 {
		return errno.New(errno.ERROR_USERNAME_USED, err)
	}
	return nil
}

// 新增用户
func CreateUser(user *User) error {
	enPw, err := encrypt.EncryptPw(user.Password)
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	user.Password = enPw

	err = db.Create(user).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 查询用户列表
func GetUsers(pageSize, pageNum int) ([]User, error) {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errno.New(errno.ERROR, err)
	}
	return users, nil
}

//编辑用户
