package middlewares

import (
	"encoding/json"
	"net/http"
	"simple-store-api/conf"
	"simple-store-api/constants"
	"simple-store-api/datatransfers"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	jwt.StandardClaims
	UID      uint   `json:"uid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type JWTConfig struct {
	JWTSignatureKey   string
	JWTPublicKey      string
	JWTExpirationTime time.Duration
}

type UserData struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
}

func VerifyToken(ctx *context.Context) {
	userData, err := doVerifyToken(ctx.Request)
	if err != nil {
		errAuth(ctx, err)
		return
	}

	ctx.Input.SetData("uid", userData.UID)
}

func doVerifyToken(r *http.Request) (result *UserData, err error) {

	token, err := getToken(r)
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	isVerified, claims, err := parseTokenJWT(token)
	if err != nil {
		err = &datatransfers.CustomError{
			Code:    constants.InternalServerErrCode,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}

	if !isVerified {
		return
	}

	result = &UserData{
		UID:      claims.UID,
		Username: claims.Username,
	}
	return

}

func getToken(r *http.Request) (token string, err error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		err = &datatransfers.CustomError{
			Code:    constants.NotAuthorizedErrCode,
			Status:  http.StatusUnauthorized,
			Message: "EMPTY_TOKEN",
		}
		return
	}

	s := strings.Split(authHeader, " ")
	if len(s) != 2 {
		err = &datatransfers.CustomError{
			Code:    constants.NotAuthorizedErrCode,
			Status:  http.StatusUnauthorized,
			Message: "INVALID_TOKEN",
		}
		return
	}

	token = s[1]
	return
}

func parseTokenJWT(token string) (isVerified bool, result *JWTClaims, err error) {
	result = &JWTClaims{}
	jwtClaims, err := jwt.ParseWithClaims(token, result, func(token *jwt.Token) (interface{}, error) {
		return conf.AppConfig.JWTConfig.JWTSignatureKey, nil
	})

	if result == nil || jwtClaims == nil || !jwtClaims.Valid || err != nil {
		return
	}
	isVerified = true
	return
}

func errAuth(ctx *context.Context, err error) {
	ctx.Output.SetStatus(http.StatusUnauthorized)
	errResponse := &datatransfers.CustomError{
		Code:    constants.NotAuthorizedErrCode,
		Status:  http.StatusUnauthorized,
		Message: err.Error(),
	}

	resBody, err := json.Marshal(errResponse)
	if err != nil {
		panic(err)
	}
	ctx.Output.Body(resBody)
}
