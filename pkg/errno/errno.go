package errno

import (
	"encoding/json"
	"time"
)

// var _ Error = (*Err)(nil)

// 正确返回
// c.JSON(http.StatusOK, errno.OK.WithData(data).WithID(c.GetString("trace-id")))

type Error interface {
	// i 为了避免被其他包实现
	i()
	WithData(data interface{}) Error
	WithID(id string) Error
	ToString() string
}

type Err struct {
	Code int         `json:"code"`           // 业务编码
	Msg  string      `json:"msg"`            // 错误描述
	Data interface{} `json:"data,omitempty"` // 返回的数据
	Time int64       `json:"time,omitempty"` // 时间戳
	ID   string      `json:"id,omitempty"`   // 当前请求的唯一ID，便于问题定位，忽略也可以
}

func NewError(code int, msg string) Error {
	return &Err{
		Code: code,
		Msg:  msg,
		Time: time.Now().Unix(),
		Data: nil,
	}
}

func (e *Err) i() {}

func (e *Err) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e *Err) WithID(id string) Error {
	e.ID = id
	return e
}

// ToString 返回 JSON 格式的错误详情
func (e *Err) ToString() string {
	Err := &Err{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
		Time: time.Now().Unix(),
		ID:   e.ID,
	}

	raw, _ := json.Marshal(Err)
	return string(raw)
}
