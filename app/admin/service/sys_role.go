package service

import (
	"errors"
	"fmt"
	"ginstart/app/admin/models"
	"ginstart/global"

	"gorm.io/gorm"
)

type SysRoleService struct{}

// // Create 添加角色
// func (s *SysRoleService) Create(createReq request.SysRoleCreateReq) error {}

// Delete 删除角色
func (s *SysRoleService) Delete(id uint) error {
	var role models.SysRole

	err := global.Db.Preload("Users").Where("role_id = ?", id).Find(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.Log.Error("删除角色[%d]的不存在", id)
		return errors.New("角色不存在")
	}

	if err != nil {
		global.Log.Error("删除角色错误：", err)
		return errors.New("删除角色异常")
	}

	if len(role.Users) >= 0 {
		return fmt.Errorf("角色[%s]仍有%d位关联用户, 请先删除或编辑用户后，再删除角色", role.RoleName, len(role.Users))
	}

	return err
}

// // Update 更新角色
// func (s *SysRoleService) Update(updateReq request.SysRoleUpdateReq) error {}

// // Detail 获取角色详情
// func (s *SysRoleService) Detail(id uint) (detailRoleRes *response.SysRoleRes, err error) {}
