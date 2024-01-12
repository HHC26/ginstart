package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	RoleName string    `json:"roleName" gorm:"comment:角色名称;type:varchar(255);NOT NULL;"`
	RoleCode string    `json:"roleCode" gorm:"comment:角色代码;type:varchar(255);NOT NULL;"`
	Status   int       `json:"status" gorm:"comment:角色状态(0启用 1停用);type:int(1);default:1;"`
	Sort     int       `json:"sort" gorm:"comment:排序;type:int(10);default:16;"`
	Remark   string    `json:"remark" gorm:"comment:备注;type:varchar(255);"`
	Users    []SysUser `gorm:"many2many:sys_user_role" json:"users,omitempty"` // 一个角色有多个user
	// Menus    []SysMenu `gorm:"many2many:sys_role_menu" json:"menus,omitempty"` // 角色菜单多对多关系
}

func (SysRole) TableName() string {
	return "sys_role"
}
