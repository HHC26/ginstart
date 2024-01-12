package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName string    `json:"userName" gorm:"comment:用户账号;type:varchar(255);NOT NULL"`
	NickName string    `json:"nickName" gorm:"comment:用户昵称;type:varchar(255);NOT NULL"`
	Email    string    `json:"email" gorm:"comment:用户邮箱;type:varchar(25);NOT NULL"`
	Phone    string    `json:"phone" gorm:"comment:手机号码;type:varchar(25);NOT NULL"`
	Password string    `json:"password" gorm:"comment:密码;type:varchar(255);NOT NULL"`
	Salt     string    `json:"salt" gorm:"comment:密码盐;type:varchar(25);"`
	Gender   int       `json:"gender" gorm:"comment:用户性别(0未知,1男,2女);type:int(1);default:0;"`
	Avatar   string    `json:"avatar" gorm:"comment:头像地址;type:varchar(255)"`
	Status   int       `json:"status" gorm:"comment:帐号状态(0正常 1停用);type:int(1);default:0;"`
	Remark   string    `json:"remark" gorm:"comment:备注;type:varchar(255)"`
	Roles    []SysRole `gorm:"many2many:sys_user_role" json:"roles,omitempty"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
