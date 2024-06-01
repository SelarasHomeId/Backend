package dto

import (
	"fmt"
	"selarashomeid/internal/config"
	"selarashomeid/internal/model"
	modeltoken "selarashomeid/internal/model/token"

	"github.com/golang-jwt/jwt"
)

type AuthLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	model.AdminEntityModel
}

type RefreshTokenRequest struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (r RefreshTokenRequest) AccessTokenClaims() (*modeltoken.AccessTokenClaims, error) {
	token, err := jwt.Parse(r.AccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(config.Get().JWT.SecretKey), nil
	})
	if token == nil || !token.Valid || err != nil {
		if jwtErrValidation, ok := err.(*jwt.ValidationError); ok {
			c := token.Claims.(jwt.MapClaims)
			return &modeltoken.AccessTokenClaims{
				ID:       c["id"].(string),
				Username: c["username"].(string),
				Email:    c["email"].(string),
				Exp:      int64(c["exp"].(float64)),
			}, jwtErrValidation
		}
		return nil, jwt.NewValidationError("invalid_access_token", jwt.ValidationErrorMalformed)
	}
	c := token.Claims.(jwt.MapClaims)
	return &modeltoken.AccessTokenClaims{
		ID:       c["id"].(string),
		Username: c["username"].(string),
		Email:    c["email"].(string),
		Exp:      int64(c["exp"].(float64)),
	}, nil
}

func (r RefreshTokenRequest) RefreshTokenClaims() (*modeltoken.RefreshTokenClaims, error) {
	token, err := jwt.Parse(r.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte(config.Get().JWT.SecretKey), nil
	})
	if token == nil || !token.Valid || err != nil {
		if jwtErrValidation, ok := err.(*jwt.ValidationError); ok {
			c := token.Claims.(jwt.MapClaims)
			return &modeltoken.RefreshTokenClaims{
				Exp: int64(c["exp"].(float64)),
			}, jwtErrValidation
		}
		return nil, jwt.NewValidationError("invalid_refresh_token", jwt.ValidationErrorMalformed)
	}
	c := token.Claims.(jwt.MapClaims)
	return &modeltoken.RefreshTokenClaims{
		Exp: int64(c["exp"].(float64)),
	}, nil
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ChangePasswordRequest struct {
	Id          int    `param:"id" validate:"required"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}
