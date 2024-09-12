package postgres

import (
	"auth/genproto/auth"
	"database/sql"
)

type AuthStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func (stg *AuthStorage) GetProfile(req *auth.ById) (*auth.GetProfileResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) UpdateProfile(req *auth.UpdateProfileRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) DeleteProfile(req *auth.ById) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) ListAllProfiles(req *auth.ById) (*auth.GetAllProfilesResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) DeleteProfileByAdmin(req *auth.DeleteProfileByAdminReq) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) RegisterUser(req *auth.Register) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) LoginUser(req *auth.Login) (*auth.TokenResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) RegisterCourier(req *auth.Register) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) LoginCourier(req *auth.Login) (*auth.TokenResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) VerifyEmail(req *auth.VerifyEmailRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) RefreshToken(req *auth.ById) (*auth.TokenResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) ChangePassword(req *auth.ChangePasswordRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) ForgetPassword(req *auth.ForgetPasswordRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) ResetPassword(req *auth.ResetPasswordRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) ChangeEmail(req *auth.ChangeEmailRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (stg *AuthStorage) VerifyNewEmail(req *auth.VerifyNewEmailRequest) (*auth.TokenResponse, error) {
	return nil, nil
}
