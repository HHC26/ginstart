package request

type ResetPwdReq struct {
	UserId   int64  `json:"userId" validate:"required,min=1" label:"用户编号"`        // 用户Id
	Password string `json:"password" validate:"required,min=8,max=50" label:"密码"` // 密码
}

type UpdatePwdReq struct {
	Password    string `json:"password" validate:"required,min=8,max=50" label:"旧密码"`    // 密码
	NewPassword string `json:"newPassword" validate:"required,min=8,max=50" label:"新密码"` // 新密码
}

type SetUserStateReq struct {
	UserId uint `json:"userId" validate:"required,min=1" label:"用户编号"` // 用户Id
	Status int  `json:"status" validate:"oneof=0 1" label:"帐号状态"`      //帐号状态（0正常 1停用）
}
