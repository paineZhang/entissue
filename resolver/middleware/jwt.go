package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/pascaldekloe/jwt"
)

const _sysKey = "LJU09hkn90inhug98721nlk;j"

const (
	userIdKey      = "userId"
	accountNameKey = "accountName"
)

func CreateJWTToken(
	userId int,
	accountName string,
	loginTime time.Time) (tokenString string, err error) {

	claims := jwt.Claims{}

	claims.ID = "com.mx.wing"
	claims.Subject = accountName
	claims.Issuer = "com.mx.wing"
	claims.Issued = jwt.NewNumericTime(loginTime.Round(time.Second))
	claims.Set = map[string]interface{}{
		userIdKey:      strconv.Itoa(userId),
		accountNameKey: accountName}

	token, err := claims.HMACSign(jwt.HS256, []byte(_sysKey))

	return string(token), err
}

func getAuthInfo(request *http.Request) (*AuthInfo, error) {

	var claims *jwt.Claims
	var err error

	claims, err = jwt.HMACCheckHeader(request, []byte(_sysKey))

	if err != nil {
		return nil, err
	}

	info := AuthInfo{}

	info.UserId, err = strconv.Atoi(claims.Set[userIdKey].(string))
	if err != nil {
		return nil, err
	}

	info.AccountName = claims.Set[accountNameKey].(string)

	return &info, nil
}
