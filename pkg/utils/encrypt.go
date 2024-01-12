package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

// MD5 generate md5 by string
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Get16BitStrMd5(str string) string {
	return MD5(str)[8:24]
}

// 编码
func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// 解码
func Base64Decode(input string) string {
	ret, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	return string(ret)
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
