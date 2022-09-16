package logic

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoCore/config"
)

// 作为生成token的参数
type OrderClaims struct {
	order *entity.Order
	//jwt-go提供的标准claim
	jwt.StandardClaims
}
type _orderToken struct {
	secret []byte
	Claims *OrderClaims
}

var OrderToken *_orderToken

func InitToken() {
	OrderToken = &_orderToken{
		secret: []byte(config.App.GetString("token.order.secret")),
	}
}

// token有效时间（纳秒）
var effectTime = 24 * time.Hour

// 生成token
func (t *_orderToken) GenerateToken(order *entity.Order) string {
	//设置token有效期
	t.Claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	t.Claims.order = order
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, t.Claims).SignedString(t.secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止
		panic(err)
	}
	return sign
}

// 解析Token
func (t *_orderToken) ParseToken(tokenString string) (*entity.Order, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &OrderClaims{}, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*OrderClaims)
	if !ok {
		return nil, errors.New("token is valid")
	}
	return claims.order, nil
}
