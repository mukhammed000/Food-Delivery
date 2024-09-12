package service

import (
	"auth/genproto/auth"
	stg "auth/storage"
	"context"
	"log"
)

type AuthService struct {
	stg stg.InitRoot
	auth.UnimplementedAuthServiceServer
}

func NewAuthService(stg stg.InitRoot) *AuthService {
	return &AuthService{
		stg: stg,
	}
}

func (s *AuthService) GetProfile(ctx context.Context, req *auth.ById) (*auth.GetProfileResponse, error) {
	resp, err := s.stg.Auth().GetProfile(req)
	if err != nil {
		log.Println("Error getting profile: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) UpdateProfile(ctx context.Context, req *auth.UpdateProfileRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().UpdateProfile(req)
	if err != nil {
		log.Println("Error updating profile: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) DeleteProfile(ctx context.Context, req *auth.ById) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().DeleteProfile(req)
	if err != nil {
		log.Println("Error deleting profile: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ListAllProfiles(ctx context.Context, req *auth.ById) (*auth.GetAllProfilesResponse, error) {
	resp, err := s.stg.Auth().ListAllProfiles(req)
	if err != nil {
		log.Println("Error listing all profiles: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) DeleteProfileByAdmin(ctx context.Context, req *auth.DeleteProfileByAdminReq) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().DeleteProfileByAdmin(req)
	if err != nil {
		log.Println("Error deleting profile by admin: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) RegisterUser(ctx context.Context, req *auth.Register) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().RegisterUser(req)
	if err != nil {
		log.Println("Error registering user: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) RegisterCourier(ctx context.Context, req *auth.Register) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().RegisterCourier(req)
	if err != nil {
		log.Println("Error registering courier: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.Login) (*auth.TokenResponse, error) {
	resp, err := s.stg.Auth().Login(req)
	if err != nil {
		log.Println("Error logging in courier: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) VerifyEmail(ctx context.Context, req *auth.VerifyEmailRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().VerifyEmail(req)
	if err != nil {
		log.Println("Error verifying email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *auth.ById) (*auth.TokenResponse, error) {
	resp, err := s.stg.Auth().RefreshToken(req)
	if err != nil {
		log.Println("Error refreshing token: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, req *auth.ChangePasswordRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().ChangePassword(req)
	if err != nil {
		log.Println("Error changing password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ForgetPassword(ctx context.Context, req *auth.ForgetPasswordRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().ForgetPassword(req)
	if err != nil {
		log.Println("Error forgetting password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().ResetPassword(req)
	if err != nil {
		log.Println("Error resetting password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ChangeEmail(ctx context.Context, req *auth.ChangeEmailRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().ChangeEmail(req)
	if err != nil {
		log.Println("Error changing email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) VerifyNewEmail(ctx context.Context, req *auth.VerifyNewEmailRequest) (*auth.TokenResponse, error) {
	resp, err := s.stg.Auth().VerifyNewEmail(req)
	if err != nil {
		log.Println("Error verifying new email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) CheckEmailAndPassword(ctx context.Context, req *auth.CheckEmailAndPasswordRequest) (*auth.InfoResponse, error) {
	resp, err := s.stg.Auth().CheckEmailAndPassword(req)
	if err != nil {
		log.Println("Error verifying new email: ", err)
		return nil, err
	}
	return resp, nil
}

