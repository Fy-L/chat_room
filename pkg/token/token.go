package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	BACKSTAGE_KEY       = "NW2eAbpR0LqUgIFwmGvykxKrohQH5Bs"
	TokenExpireDuration = time.Hour * 24
	// TokenExpireDuration = time.Second * 5
)

type CustomClaims struct {
	Rtm string `json:"rtm"`
	Iat int64  `json:"iat"`
	Exp int64  `json:"exp"`
	Ctm string `json:"ctm"`
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

// 生成token
func GenToken(uid string) (string, error) {
	iat := time.Now()
	exp := iat.Add(TokenExpireDuration).Unix()
	var c = CustomClaims{
		Rtm: "test",
		Iat: iat.Unix(),
		Exp: exp,
		Uid: uid,
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(BACKSTAGE_KEY))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(BACKSTAGE_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
