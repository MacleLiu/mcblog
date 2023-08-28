package models

import (
	"errors"
	"mcblog/utils/encrypt"
	"mcblog/utils/errno"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(200);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int" json:"role" validate:"required,gte=2" label:"角色码"`
}

// 检查用户是否存在
func CheckUser(id int) error {
	var user User
	err := db.Select("username").Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_USER_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	return nil
}

// 检查用户名是否已使用
func CheckUserName(name string) error {
	var user User
	err := db.Select("id").Where("username=?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_USER_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}
	if user.ID > 0 {
		return nil
	}
	return errno.New(errno.ERROR, err)
}

// 检查新用户名是否被其他用户使用
func CheckUpUser(id int, name string) error {
	var user User
	err := db.Select("id").Where("username=?", name).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_USER_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}
	if user.ID == uint(id) {
		return nil
	}
	return errno.New(errno.ERROR_USERNAME_USED, err)
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

// 查询单个用户
func GetUser(id int) (User, error) {
	var user User
	err := db.Select("username", "role").Where("id = ?", id).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errno.New(errno.ERROR_USER_NOT_EXIST, err)
		}
		return user, errno.New(errno.ERROR, err)
	}
	return user, nil
}

// 查询用户列表
func GetUsers(pageSize, pageNum int, username string) ([]User, int64, error) {
	var users []User
	var total int64
	var err error
	if username == "" {
		err = db.Model(&users).Select("id", "username", "role").Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	} else {
		err = db.Select("id", "username", "role").Where("username LIKE ?", username+"%").Find(&users).Count(&total).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return users, total, nil
}

// 编辑用户信息
func EditUser(id int, data User) error {
	var user User
	dataMap := make(map[string]any)
	dataMap["username"] = data.Username
	dataMap["role"] = data.Role

	err = db.Model(&user).Where("id = ?", id).Updates(dataMap).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 删除用户
func DeleteUser(id int) error {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 登陆验证
func LoginVerify(username, password string) error {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_USER_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	// 未做精细化权限管理，暂时禁止非管理员用户登录
	if user.Role >= 2 {
		return errno.New(errno.ERROR_AUTH_DEFICIENT, err)
	}

	// 比对密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errno.New(errno.ERROR_PASSWORD_WRONG, err)
	}

	return nil
}
