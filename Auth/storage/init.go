package storage

import (
	auth "auth/genproto/auth"
)

type InitRoot interface {
	Auth() AuthService
}

type AuthService interface {
	// Profile Management
	GetProfile(req *auth.ById) (*auth.GetProfileResponse, error)
	UpdateProfile(req *auth.UpdateProfileRequest) (*auth.InfoResponse, error)
	DeleteProfile(req *auth.ById) (*auth.InfoResponse, error)

	// Admin Operations
	ListAllProfiles(req *auth.ById) (*auth.GetAllProfilesResponse, error)
	DeleteProfileByAdmin(req *auth.DeleteProfileByAdminReq) (*auth.InfoResponse, error)

	// Users and Couriers
	RegisterUser(req *auth.Register) (*auth.InfoResponse, error)
	RegisterCourier(req *auth.Register) (*auth.InfoResponse, error)
	Login(req *auth.Login) (*auth.TokenResponse, error)

	// Both
	VerifyEmail(req *auth.VerifyEmailRequest) (*auth.InfoResponse, error)
	RefreshToken(req *auth.ById) (*auth.TokenResponse, error)
	ChangePassword(req *auth.ChangePasswordRequest) (*auth.InfoResponse, error)
	ForgetPassword(req *auth.ForgetPasswordRequest) (*auth.InfoResponse, error)
	ResetPassword(req *auth.ResetPasswordRequest) (*auth.InfoResponse, error)
	ChangeEmail(req *auth.ChangeEmailRequest) (*auth.InfoResponse, error)
	VerifyNewEmail(req *auth.VerifyNewEmailRequest) (*auth.TokenResponse, error)
	CheckEmailAndPassword(req *auth.CheckEmailAndPasswordRequest) (*auth.InfoResponse, error)
}
