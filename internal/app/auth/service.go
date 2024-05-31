package auth

import (
	"errors"
	"fmt"
	"net/http"
	"selarashomeid/internal/abstraction"
	"selarashomeid/internal/config"
	"selarashomeid/internal/dto"
	"selarashomeid/internal/factory"
	"selarashomeid/internal/model"
	modeltoken "selarashomeid/internal/model/token"
	"selarashomeid/internal/repository"
	"selarashomeid/pkg/util/aescrypt"
	"selarashomeid/pkg/util/encoding"
	"selarashomeid/pkg/util/response"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error)
	Logout(ctx *abstraction.Context) (map[string]interface{}, error)
	RefreshToken(ctx *abstraction.Context, payload *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
}

type service struct {
	AdminRepository repository.Admin

	DB *gorm.DB
}

func NewService(f *factory.Factory) Service {
	return &service{
		AdminRepository: f.AdminRepository,

		DB: f.Db,
	}
}

func (s *service) encryptTokenClaims(v int) (encryptedString string, err error) {
	encryptedString, err = aescrypt.EncryptAES(fmt.Sprint(v), config.Get().JWT.SecretKey)
	return
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.AuthLoginRequest) (*dto.AuthLoginResponse, error) {
	data, err := s.AdminRepository.FindByUsername(ctx, payload.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.Unauthorized, errors.New("username or password is incorrect"))
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password)); err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.Unauthorized, errors.New("username or password is incorrect"))
	}

	var encryptedUserID string
	if encryptedUserID, err = s.encryptTokenClaims(data.ID); err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	encodedUsername := encoding.Encode(data.Username)
	encodedEmail := encoding.Encode(data.Email)

	accessTokenClaims := &modeltoken.AccessTokenClaims{
		ID:       encryptedUserID,
		Username: encodedUsername,
		Email:    encodedEmail,
		Exp:      time.Now().Add(time.Duration(1 * time.Hour)).Unix(),
	}
	authToken := modeltoken.NewAuthToken(accessTokenClaims)
	accessToken, err := authToken.AccessToken()
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}
	refreshToken, err := authToken.RefreshToken()
	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	return &dto.AuthLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AdminEntityModel: model.AdminEntityModel{
			ID: data.ID,
			AdminEntity: model.AdminEntity{
				Name:     data.Name,
				Email:    data.Email,
				Username: data.Username,
			},
		},
	}, nil
}

func (s *service) Logout(ctx *abstraction.Context) (map[string]interface{}, error) {
	return map[string]interface{}{
		"message": "Logout successful",
	}, nil
}

func (s *service) RefreshToken(ctx *abstraction.Context, payload *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	accessTokenClaims, err := payload.AccessTokenClaims()
	if err != nil && err.(*jwt.ValidationError).Errors != jwt.ValidationErrorExpired {
		return nil, response.CustomErrorBuilder(http.StatusBadRequest, "invalid_access_token", "invalid_access_token", err.Error())
	}
	accessTokenAuthCtx, err := accessTokenClaims.AuthContext()
	if err != nil {
		return nil, response.CustomErrorBuilder(http.StatusBadRequest, err.Error(), "invalid_access_token", err.Error())
	}

	refreshTokenClaims, err := payload.RefreshTokenClaims()
	if err != nil {
		if jwtValErr := err.(*jwt.ValidationError); jwtValErr.Errors == jwt.ValidationErrorExpired {
			return nil, response.CustomErrorBuilder(http.StatusUnauthorized, "refresh_token_is_expired", "refresh_token_is_expired", err.Error())
		} else {
			return nil, response.CustomErrorBuilder(http.StatusBadRequest, jwtValErr.Error(), "invalid_refresh_token", err.Error())
		}
	}
	refreshTokenAuthCtx, err := refreshTokenClaims.AuthContext()
	if err != nil {
		return nil, response.CustomErrorBuilder(http.StatusBadRequest, err.Error(), "invalid_refresh_token", err.Error())
	}

	if refreshTokenAuthCtx.ID != accessTokenAuthCtx.ID {
		return nil, response.CustomErrorBuilder(http.StatusUnauthorized, "unauthorized_to_refresh_token", "unauthorized_to_refresh_token", "access token cannot contains with refresh token")
	}

	data, err := s.AdminRepository.FindById(ctx, refreshTokenAuthCtx.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.Unauthorized, errors.New("username or password is incorrect"))
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	accessTokenClaims = refreshTokenClaims.AccessTokenClaims()
	accessTokenClaims.Username = encoding.Encode(data.Username)
	accessTokenClaims.Email = encoding.Encode(data.Email)

	authToken := modeltoken.NewAuthToken(accessTokenClaims)
	accessToken, err := authToken.AccessToken()
	if err != nil {
		return nil, response.CustomErrorBuilder(http.StatusUnauthorized, err.Error(), "err_generate_access_token", err.Error())
	}
	refreshToken, err := authToken.RefreshToken()
	if err != nil {
		return nil, response.CustomErrorBuilder(http.StatusUnauthorized, err.Error(), "err_generate_refresh_token", err.Error())
	}

	return &dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
