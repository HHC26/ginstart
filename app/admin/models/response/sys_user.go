package response

import (
	"time"
)

// SysUserRes 输出对象
type SysUserRes struct {
	Id        uint      `json:"id"`                         //编号
	UserName  string    `json:"userName"`                   //用户账号
	NickName  string    `json:"nickName" `                  //用户昵称
	Email     string    `json:"email" `                     //用户邮箱
	Phone     string    `json:"phone"`                      //手机号码
	Gender    int       `json:"gender"`                     //用户性别（0未知，1男，2女）
	Avatar    string    `json:"avatar"`                     //头像地址
	Status    int       `json:"status"`                     //帐号状态（0正常 1停用）
	Remark    string    `json:"remark"`                     //备注
	CreatedAt time.Time `json:"createdAt"`                  //创建时间
	UpdatedAt time.Time `json:"updatedAt"`                  //更新时间
	RoleIds   []uint    `json:"roleIds,omitempty" gorm:"-"` //角色
}
