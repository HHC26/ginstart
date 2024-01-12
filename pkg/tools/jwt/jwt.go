package jwt

import (
	"errors"
	"ginstart/global"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserAuthClaims struct {
	UserName string
	NickName string
	jwt.RegisteredClaims
}

// 用于加盐的字符串
var secretKey []byte

// GenerateToken 生成token
func GenerateToken(claims *UserAuthClaims) (string, error) {
	j := global.Conf.Jwt
	var jwtTimeout = j.ExpiresTime
	var verifyKey = j.SigningKey
	var issuer = j.Issuer

	secretKey = []byte(verifyKey + "RcWJhezAYXjzVzZRmaX8B8uy2U3Hgg2p")

	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(jwtTimeout))), // 过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                                              // 签发时间
		NotBefore: jwt.NewNumericDate(time.Now()),                                              // 生效时间
		Issuer:    issuer,                                                                      // 签发人
		Subject:   "your_admin",                                                                // 签名主题
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secretKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*UserAuthClaims, error) {
	// 解析 token string 拿到 token jwt.Token
	var claim UserAuthClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*UserAuthClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GetUserInfo 获取当前用户信息
func GetUserInfo(c *gin.Context) *UserAuthClaims {
	claims, _ := c.Get("claims")
	userInfo := claims.(*UserAuthClaims)
	return userInfo
}

// GetUserId 获取当前用户名
func GetUserName(c *gin.Context) string {
	userInfo := GetUserInfo(c)
	return userInfo.UserName
}
