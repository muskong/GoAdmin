package logic

import (
	"errors"

	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoPkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type (
	LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	JwtData struct {
		Name  string
		Token string
	}
)

func LoginVerify(data LoginData) (jwtData JwtData, err error) {

	userData, err := entity.GetAdminName(data.Username)
	if userData.Id <= 0 || err != nil {
		err = errors.New("用户或密码出错1")
		return
	}

	is := checkHashPassword(data.Password, userData.Password)
	if !is {
		err = errors.New("用户或密码出错2")
		return
	}

	jwtData.Name = userData.Name
	jwtData.Token = jwt.Jwt.GenerateToken(userData.Id)

	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
